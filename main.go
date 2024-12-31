package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// Post represents a single blog post.
type Post struct {
	ID      int
	Title   string
	Content string
}

// Sample posts stored in memory.
var posts = []Post{
	{ID: 1, Title: "Welcome to My Blog", Content: "This is my first blog post."},
	{ID: 2, Title: "Learning Go", Content: "Go is a fantastic language."},
	{ID: 3, Title: "Dockerizing Applications", Content: "Docker makes deployment so much easier."},
}

// Handler to display the homepage with a list of posts.
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/index.html")
	if err != nil {
		http.Error(w, "Error loading templates :(", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Pass the list of posts to the template.
	err = tmpl.ExecuteTemplate(w, "layout", posts)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Println(err)
	}
}

// Handler to display a single post based on ID
func PostHandler(w http.ResponseWriter, r *http.Request) {
	// Get the "id" query parameter.
	id := r.URL.Query().Get("id")

	// TODO: I just wanna see what's the id when we click on the link.
	fmt.Println(id)

	// Find the post with the matching ID.
	for _, post := range posts {
		// TODO: Now, show me the post id:
		fmt.Println(post.ID)

		if id == string(rune(post.ID)) {

			tmpl, err := template.ParseFiles("templates/layout.html", "templates/post.html")
			if err != nil {
				http.Error(w, "Error loading templates", http.StatusInternalServerError)
				log.Println(err)
				return
			}

			// Pass the post to the template.
			err = tmpl.ExecuteTemplate(w, "layout", post)
			if err != nil {
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
				log.Println(err)
			}
			return
		}
	}

	// If no post is found, return a 404.
	http.NotFound(w, r)
}

func main() {
	// Routes
	http.HandleFunc("/", HomePageHandler)
	http.HandleFunc("/post", PostHandler)

	// Start the server.
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
