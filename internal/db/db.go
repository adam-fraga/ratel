/*
Copyright © 2024 Adm FRG adam.fraga@live.fr
Package db hold the struct and logic to interact with the database system of the ratel web framework.
*/

package db

import (
	"database/sql"
	"os"

	"github.com/adam-fraga/ratel/internal/errors"
	"github.com/adam-fraga/ratel/utils"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Db represent the db session and hold a database connection
type Db struct {
	Conn *sql.DB
}

// Connect Create a new db session and connect to the database
func (pg *Db) Connect() error {
	err := godotenv.Load()
	if err != nil {
		return &errors.DevError{
			Type:       "ENV",
			Msg:        "Error loading .env file: " + err.Error(),
			Origin:     "db.Connect()",
			FileOrigin: "internal/db/db.go"}
	}
	var user, dbname, password string = os.Getenv("RATEL_DB_USER"), os.Getenv("RATEL_DB_NAME"), os.Getenv("RATEL_DB_PASSWORD")
	connStr := "user=" + user + " dbname=" + dbname + " password=" + password + " sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return &errors.DbError{
			Query:  connStr,
			Msg:    err.Error(),
			Action: "sql.Open()"}
	}
	pg.Conn = db

	utils.PrintInfoMsg("Connected to the database")
	return nil
}

// Close the database connection
func (pg *Db) Close() {
	pg.Conn.Close()
	utils.PrintInfoMsg("Database connection closed")
}
