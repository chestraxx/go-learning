package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createTableUser := `
	CREATE TABLE IF NOT EXISTS user (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);`
	_, err := DB.Exec(createTableUser)
	if err != nil {
		panic(err)
	}

	createTableEvent := `
	CREATE TABLE IF NOT EXISTS event (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY(user_id) REFERENCES user(id)
	);`
	_, err = DB.Exec(createTableEvent)
	if err != nil {
		panic(err)
	}

	createTableRegistration := `
	CREATE TABLE IF NOT EXISTS registration (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY(event_id) REFERENCES event(id),
		FOREIGN KEY(user_id) REFERENCES user(id)
	);`
	_, err = DB.Exec(createTableRegistration)
	if err != nil {
		panic(err)
	}
}
