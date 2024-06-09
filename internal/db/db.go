/*
Copyright © 2024 Adm FRG adam.fraga@live.fr
Package db hold the struct and logic to interact with the database system of the ratel web framework.
*/

// This package contains the struct and logic to interact with the database system of the ratel web framework.
package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Db represent the db session and hold a database connection
type Db struct {
	Conn   *sql.DB
	tables []string
}

// NewDb create a new db session
func NewDb() (*Db, error) {
	conn, err := sql.Open("sqlite3", "ratel.db")

	if err != nil {
		return nil, err
	}
	return &Db{
		Conn: conn,
	}, nil
}

// Close close the db session
func (db *Db) Close() error {
	return db.Conn.Close()
}

// InitDatabase create the tables in the database
func (db *Db) InitDatabase(table Table) error {

	for _, table := range db.tables {
		if err := db.createTable(table); err != nil {
			return err
		}
	}
	return nil
}

// Init create the tables in the database
func (db *Db) createTable(table Table) error {
	var fields string
	for _, field := range table.Fields {
		fields += field.Name + " " + field.Type
		if field.NotNull {
			fields += " NOT NULL"
		}
		if field.Unique {
			fields += " UNIQUE"
		}
		if field.Default != "" {
			fields += " DEFAULT " + field.Default
		}
		fields += ", "
	}
	fields = fields[:len(fields)-2]

	_, err := db.Conn.Exec("CREATE TABLE IF NOT EXISTS " + table.Name + " (" + fields + ")")
	if err != nil {
		return err
	}
	return nil
}
