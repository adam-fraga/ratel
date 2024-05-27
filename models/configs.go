package models

// DbUserConfig is a struct that holds the database configuration for the command dbCreateContainer
type DbUserConfig struct {
	DbProvider string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}
