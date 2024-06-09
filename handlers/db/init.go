/*
Copyright © 2024 Adm FRG adam.fraga@live.fr
Package db contains handlers to execute the logic of the database system of ratel web framework
*/

package db

import (
	"github.com/adam-fraga/ratel/internal/errors"
)

// InitDb initialize the database for ratel
func InitDb() error {
	if err := createDatabase(); err != nil {
		return &errors.DevError{
			Type:       "createDatabase()",
			Origin:     "InitDb()",
			FileOrigin: "init.go",
			Msg:        "Error creating the database: " + err.Error(),
		}
	}

	return nil
}

// InitSchema initialize the database schema for the project using ratel web framework
func initSchema() error {
	return nil
}

// createDatabase create the database for the necessary data for the project using ratel web framework
func createDatabase() error {
	return nil
}
