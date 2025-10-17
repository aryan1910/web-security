package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type Server struct {
	mu       sync.Mutex
	comments []template.HTML
	tmpl     *template.Template
}

func main() {

	tmplPath := filepath.Join("templates", "index.html")
	tmpl := template.Must(template.ParseFiles(tmplPath))

	server := &Server{
		tmpl:     tmpl,
		comments: []template.HTML{},
	}

	http.HandleFunc("/", server.handleComments)

	log.Println("⚠️  XSS demo server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (s *Server) handleComments(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Invalid form submission", http.StatusBadRequest)
			return
		}

		comment := r.FormValue("comment")

		s.mu.Lock()
		s.comments = append(s.comments, template.HTML(comment))
		s.mu.Unlock()

		http.Redirect(w, r, "/", http.StatusSeeOther)

	case http.MethodGet:
		s.mu.Lock()
		data := struct {
			Comments []template.HTML
		}{
			Comments: s.comments,
		}
		s.mu.Unlock()

		if err := s.tmpl.Execute(w, data); err != nil {
			http.Error(w, "Template rendering error", http.StatusInternalServerError)
		}

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
