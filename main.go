package main

import (
	"blog/database"
	"blog/handlers"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Initialize the database.
	database.Init()

	// Routes
	http.HandleFunc("/", handlers.HomePageHandler)
	http.HandleFunc("/post", handlers.DisplayPostHandler)
	http.HandleFunc("/post/add", handlers.AddPostHandler)

	// Start the server.
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
