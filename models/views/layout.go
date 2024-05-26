package views

import (
	"database/sql"
	"github.com/adam-fraga/ratel/internal/db"
	"github.com/adam-fraga/ratel/internal/errors"
	"github.com/adam-fraga/ratel/utils"
)

type Layout struct {
	Name        string
	Description string
	pages       []Page
	Path        string `default:"views/layouts"`
}

func (l *Layout) Create() error {

	var db db.Db

	err := db.Connect()
	defer db.Close()

	if err != nil {
		return &errors.DevError{Msg: "Error connecting to the database: " + err.Error()}
	}

	_, err2 := db.Conn.Exec(`INSERT INTO layouts (name, description) VALUES ($1, $2, $3)`, l.Name, l.Description, l.pages)

	if err2 != nil {
		return &errors.DbError{Query: "INSERT INTO layouts", Msg: err2.Error(), Action: "db.Exec()"}
	}

	utils.PrintSuccessMsg("Layout created !")

	return nil
}

func (l *Layout) Delete() error {

	var db db.Db

	err := db.Connect()
	defer db.Close()

	if err != nil {
		return &errors.DevError{Msg: "Error connecting to the database: " + err.Error()}
	}

	_, err2 := db.Conn.Exec(`DELETE FROM layouts WHERE name=$1`, l.Name)

	if err2 != nil {
		return &errors.DbError{Query: "DELETE FROM layouts", Msg: err2.Error(), Action: "db.Exec()"}
	}

	utils.PrintSuccessMsg("Layout deleted !")

	return nil
}

func (l *Layout) Update() error {

	var db db.Db

	err := db.Connect()
	defer db.Close()

	if err != nil {
		return &errors.DevError{Msg: "Error connecting to the database: " + err.Error()}
	}

	_, err2 := db.Conn.Exec(`UPDATE layouts SET name=$1, description=$2, pages=$3 WHERE name=$4`, l.Name, l.Description, l.pages, l.Name)

	if err2 != nil {
		return &errors.DbError{Query: "UPDATE layouts", Msg: err2.Error(), Action: "db.Exec()"}
	}

	utils.PrintSuccessMsg("Layout updated !")

	return nil
}

func (l *Layout) Get() error {
	var db db.Db

	err := db.Connect()
	defer db.Close()

	if err != nil {
		return &errors.DevError{Msg: "Error connecting to the database: " + err.Error()}
	}

	rows, err2 := db.Conn.Query(`SELECT * FROM layouts WHERE name=$1`, l.Name)

	if err2 != nil {
		return &errors.DbError{Query: "SELECT * FROM layouts", Msg: err2.Error(), Action: "db.Query()"}
	}

	for rows.Next() {
		err := rows.Scan(&l.Name, &l.Description, &l.pages)
		if err != nil {
			return &errors.DbError{Query: "rows.Scan()", Msg: err.Error(), Action: "rows.Scan()"}
		}
	}

	utils.PrintSuccessMsg("Layout retrieved !")

	return nil
}

//SCHEMAS

func CreateLayoutTable(conn *sql.DB) error {

	_, err2 := conn.Exec(` CREATE TABLE IF NOT EXISTS layouts (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT
  )`)

	if err2 != nil {
		return &errors.DbError{Query: "CREATE TABLE layouts", Msg: err2.Error(), Action: "db.Exec()"}
	}

	utils.PrintSuccessMsg("Table layouts created !")

	return nil
}

func DeleteLayoutTable(conn *sql.DB) error {

	_, err2 := conn.Exec(`DROP TABLE IF EXISTS layouts`)

	if err2 != nil {
		return &errors.DbError{Query: "DROP TABLE layouts", Msg: err2.Error(), Action: "db.Exec()"}
	}

	utils.PrintSuccessMsg("Table layouts deleted !")

	return nil
}
