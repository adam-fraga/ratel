package views

import (
	"database/sql"

	"github.com/adam-fraga/ratel/internal/db"
	"github.com/adam-fraga/ratel/internal/errors"
	"github.com/adam-fraga/ratel/utils"
)

type Page struct {
	Name        string
	Description string
	Path        string `default:"views/pages"`
}

func (p *Page) Create() error {

	var db db.Db

	err := db.Connect()
	defer db.Close()

	if err != nil {
		return &errors.DevError{Msg: "Error connecting to the database: " + err.Error()}
	}

	_, err2 := db.Conn.Exec(`INSERT INTO pages (name, description) VALUES ($1, $2)`, p.Name, p.Description)

	if err2 != nil {
		return &errors.DbError{Query: "INSERT INTO pages", Msg: err2.Error(), Action: "db.Exec()"}
	}

	utils.PrintSuccessMsg("Page created !")

	return nil
}

func (p *Page) Delete() error {

	var db db.Db

	err := db.Connect()
	defer db.Close()

	if err != nil {
		return &errors.DevError{Msg: "Error connecting to the database: " + err.Error()}
	}

	_, err2 := db.Conn.Exec(`DELETE FROM pages WHERE name=$1`, p.Name)

	if err2 != nil {
		return &errors.DbError{Query: "DELETE FROM pages", Msg: err2.Error(), Action: "db.Exec()"}
	}

	utils.PrintSuccessMsg("Page deleted !")

	return nil
}

func (p *Page) Update() error {

	var db db.Db

	err := db.Connect()
	defer db.Close()

	if err != nil {
		return &errors.DevError{Msg: "Error connecting to the database: " + err.Error()}
	}

	_, err2 := db.Conn.Exec(`UPDATE pages SET name=$1, description=$2 WHERE name=$3`, p.Name, p.Description, p.Name)

	if err2 != nil {
		return &errors.DbError{Query: "UPDATE pages", Msg: err2.Error(), Action: "db.Exec()"}
	}

	utils.PrintSuccessMsg("Page updated !")

	return nil
}

func (p *Page) Read() error {

	var db db.Db

	err := db.Connect()
	defer db.Close()

	if err != nil {
		return &errors.DevError{Msg: "Error connecting to the database: " + err.Error()}
	}

	rows, err2 := db.Conn.Query(`SELECT * FROM pages WHERE name=$1`, p.Name)

	if err2 != nil {
		return &errors.DbError{Query: "SELECT * FROM pages", Msg: err2.Error(), Action: "db.Query()"}
	}

	for rows.Next() {
		err := rows.Scan(&p.Name, &p.Description)
		if err != nil {
			return &errors.DbError{Query: "rows.Scan()", Msg: err.Error(), Action: "rows.Scan()"}
		}
	}

	return nil
}

// Schema
func CreatePageTable(conn *sql.DB) error {

	_, err2 := conn.Exec(` CREATE TABLE IF NOT EXISTS pages (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT
    FOREIGN KEY (layout_id) REFERENCES layouts(id)
  )`)

	if err2 != nil {
		return &errors.DbError{Query: "CREATE TABLE layouts", Msg: err2.Error(), Action: "db.Exec()"}
	}

	utils.PrintSuccessMsg("Table layouts created !")

	return nil
}

func DeletePageTable(conn *sql.DB) error {

	_, err2 := conn.Exec(`DROP TABLE IF EXISTS pages`)

	if err2 != nil {
		return &errors.DbError{Query: "DROP TABLE pages", Msg: err2.Error(), Action: "db.Exec()"}
	}

	utils.PrintSuccessMsg("Table pages deleted !")

	return nil
}
