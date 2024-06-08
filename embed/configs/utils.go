package utils

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"path"

	er "github.com/adam-fraga/ratel/internal/errors"
)

var Messages map[string]map[string]string

// LoadMessages loads the messages from the locales directory
func LoadMessages() error {

	Messages = make(map[string]map[string]string)
	rootProject := os.Getenv("PROJECT_ROOT_PATH")
	locales, err := path.Join(rootProject, "/locales")
	localeFiles, err := locales.ReadDir(locales)
	if err != nil {
		return &er.DevError{
			Type:       "Error",
			Origin:     "LoadMessages",
			FileOrigin: "messageLoader.go",
			Msg:        fmt.Sprintf("Error Loaded content snippets: %v", err)}
	}

	for _, localeFile := range localeFiles {
		locale := localeFile.Name()
		file, err := locales.ReadFile(fmt.Sprintf("/home/afraga/Projects/ratel/utils/locales/%s", locale))
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

// GetMessage returns the message for the given language and key
func GetMessage(lang, key string) string {
	if msg, ok := Messages[lang][key]; ok {
		return msg
	}
	return key // return the key itself if not found
}
