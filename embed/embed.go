package embed

import (
	"embed"
)

//go:embed configs/*
var EmbeddedConfigs embed.FS

//go:embed projectStruct.json
var EmbeddedProjectStruct []byte

//go:embed ratel.db
var EmbeddedRatelDB []byte
