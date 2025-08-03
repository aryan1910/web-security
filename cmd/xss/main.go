package main

import (
	"html/template"
	"log"
	"net/http"
	"sync"
)

type Server struct {
	comments []template.HTML // Unsafe on purpose!
	mu       sync.Mutex
}

func main() {
	server := &Server{
		comments: []template.HTML{},
	}

	http.HandleFunc("/", server.handle)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (s *Server) handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		comment := r.FormValue("comment")

		s.mu.Lock()
		s.comments = append(s.comments, template.HTML(comment))
		s.mu.Unlock()

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tmpl := `
		<!DOCTYPE html>
		<html>
		<head><title>XSS Demo</title></head>
		<body>
			<h1>XSS Demo Page</h1>
			<form method="POST" action="/">
				<textarea name="comment" rows="4" cols="50"></textarea><br>
				<input type="submit" value="Post Comment">
			</form>

			<h3>Comments:</h3>
			{{range .}}
				<p>{{.}}</p>
			{{end}}
		</body>
		</html>
	`

	t := template.Must(template.New("page").Parse(tmpl))

	s.mu.Lock()
	defer s.mu.Unlock()
	t.Execute(w, s.comments)
}
