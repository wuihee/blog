package handlers

import (
	"blog/database"
	"blog/models"
	"database/sql"
	"html/template"
	"log"
	"net/http"
)

// Handler to display the homepage with a list of posts.
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, title, content FROM posts")
	if err != nil {
		http.Error(w, "Error fetching posts", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content)
		if err != nil {
			log.Println(err)
			continue
		}
		posts = append(posts, post)
	}

	tmpl, err := template.ParseFiles("templates/layout.html", "templates/index.html")
	if err != nil {
		http.Error(w, "Error loading templates", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	tmpl.ExecuteTemplate(w, "layout", posts)
}

// Handler to display a single post based on ID
func DisplayPostHandler(w http.ResponseWriter, r *http.Request) {
	// Get the "id" query parameter.
	id := r.URL.Query().Get("id")
	var post models.Post

	err := database.DB.QueryRow("SELECT id, title, content FROM posts WHERE id = ?", id).Scan(&post.ID, &post.Title, &post.Content)
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	} else if err != nil {
		http.Error(w, "Error fetching post", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	tmpl, err := template.ParseFiles("templates/layout.html", "templates/post.html")
	if err != nil {
		http.Error(w, "Error loading templates.", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	tmpl.ExecuteTemplate(w, "layout", post)
}

// Handles adding a new post (for testing).
func AddPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/layout.html", "templates/add.html")
		if err != nil {
			http.Error(w, "Error loading templates", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		tmpl.ExecuteTemplate(w, "layout", nil)
		return
	}

	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		content := r.FormValue("content")

		_, err := database.DB.Exec("INSERT INTO posts (title, content) VALUES (?, ?)", title, content)
		if err != nil {
			http.Error(w, "Error saving post", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
