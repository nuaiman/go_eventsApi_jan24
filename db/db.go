package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("sqlite", "api.db")
	if err != nil {
		panic("Database could not connect: " + err.Error())
	}
	DB = db
	err = createTables()
	if err != nil {
		panic("Database could not connect: " + err.Error())
	}
	fmt.Println("Tables created successfully!")
}

func createTables() error {
	createUsersTable := `
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            email TEXT NOT NULL UNIQUE,
            password TEXT NOT NULL
        )`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Couldnot create users table")
	}

	createEventsTable := `
        CREATE TABLE IF NOT EXISTS events (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
			userId INTEGER,
            name TEXT NOT NULL,
            description TEXT NOT NULL,
            location TEXT NOT NULL,
            dateTime DATETIME NOT NULL,
			FOREIGN KEY (userId) REFERENCES users (id)
        )`
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Couldnot create users table")
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		userId INTEGER,
		eventId INTEGER,
		FOREIGN KEY (userId) REFERENCES users (id),
		FOREIGN KEY (eventId) REFERENCES events (id)
	)`
	_, err = DB.Exec(createRegistrationsTable)

	return err
}
