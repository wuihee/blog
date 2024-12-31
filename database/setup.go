package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Initializes the database connection and creates necessary tables.
func Init() {
	var err error

	// Connect to SQLite database (creates file if doesn't exist).
	DB, err = sql.Open("sqlite3", "./blog.db")
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Create the posts table.
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL
	);
	`

	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	log.Println("Database initialized successfully.")
}
