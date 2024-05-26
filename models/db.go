package models

import (
	"database/sql"
	"os"

	"github.com/adam-fraga/ratel/errors"
	"github.com/adam-fraga/ratel/utils"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type PgDb struct {
	Db *sql.DB
}

func (pg *PgDb) Connect() error {
	err := godotenv.Load()
	if err != nil {
		return &errors.ClientError{Msg: "Error loading .env file: " + err.Error()}
	}
	var user, dbname, password string = os.Getenv("RATEL_DB_USER"), os.Getenv("RATEL_DB_NAME"), os.Getenv("RATEL_DB_PASSWORD")
	connStr := "user=" + user + " dbname=" + dbname + " password=" + password + " sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return &errors.ClientError{Msg: "Error connecting to the database: " + err.Error()}
	}
	pg.Db = db

	utils.PrintInfoMsg("Connected to the database")
	return nil
}

func (pg *PgDb) Close() {
	pg.Db.Close()
	utils.PrintInfoMsg("Database connection closed")
}
