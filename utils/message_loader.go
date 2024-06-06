package utils

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"

	er "github.com/adam-fraga/ratel/internal/errors"
)

//go:embed locales/*
var locales embed.FS

var Messages map[string]map[string]string

func LoadMessages() error {
	Messages = make(map[string]map[string]string)
	localeFiles, err := locales.ReadDir("github.com/adam-fraga/ratel/locales")
	if err != nil {
		return &er.DevError{
			Type:       "Error",
			Origin:     "LoadMessages",
			FileOrigin: "messageLoader.go",
			Msg:        fmt.Sprintf("Error Loaded content snippets: %v", err)}
	}

	for _, localeFile := range localeFiles {
		locale := localeFile.Name()
		file, err := locales.ReadFile(fmt.Sprintf("github.com/adam-fraga/ratel/locales/%s/messages.json", locale))
		if err != nil {
			log.Fatalf("Error reading %s: %v", locale, err)
		}

		var msgMap map[string]string
		if err := json.Unmarshal(file, &msgMap); err != nil {
			log.Fatalf("Error unmarshalling %s: %v", locale, err)
		}

		Messages[locale] = msgMap
	}
	return nil
}

func GetMessage(lang, key string) string {
	if msg, ok := Messages[lang][key]; ok {
		return msg
	}
	return key // return the key itself if not found
}
