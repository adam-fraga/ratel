package embed

import (
	"embed"
)

//go:embed configs/*
var EmbeddedConfigs embed.FS

//go:embed projectStruct.json
var EmbeddedProjectStruct []byte
