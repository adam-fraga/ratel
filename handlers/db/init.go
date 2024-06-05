/*
Copyright © 2024 Adm FRG adam.fraga@live.fr
Package db contains handlers to execute the logic of the database system of ratel web framework
*/

package db

import (
	"github.com/adam-fraga/ratel/internal/db"
	"github.com/adam-fraga/ratel/internal/errors"
	"github.com/adam-fraga/ratel/utils"
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

// InitSchema initialize the database schema for ratel web framework
func initSchema() error {
	var db db.Db

	if err := db.Connect(); err != nil {
		return &errors.DevError{Msg: "Error connecting to the database: " + err.Error()}
	}

	defer db.Close()

	return nil
}

// createDatabase create the database for the necessary data for ratel web framework
func createDatabase() error {
	var db db.Db

	err := db.Connect()
	defer db.Close()

	if err != nil {
		return &errors.DevError{Msg: "Error connecting to the database: " + err.Error()}
	}

	_, err2 := db.Conn.Exec(`CREATE DATABASE IF NOT EXISTS ratel`)
	if err2 != nil {
		return &errors.DbError{Query: "CREATE DATABASE ratel", Msg: err2.Error(), Action: "db.Exec()"}
	}

	utils.PrintSuccessMsg("Database created !")

	return nil
}
