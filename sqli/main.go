package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Setup DB
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create users table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		password TEXT
	);`)
	if err != nil {
		log.Fatal(err)
	}

	// Insert test user
	db.Exec("INSERT INTO users (username, password) VALUES ('admin', 'password123')")

	// Login endpoint (vulnerable)
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		password := r.URL.Query().Get("password")

		query := fmt.Sprintf("SELECT * FROM users WHERE username = '%s' AND password = '%s'", username, password)
		fmt.Println("Executing query:", query)

		row := db.QueryRow(query)
		var id int
		var user, pass string
		err := row.Scan(&id, &user, &pass)
		if err != nil {
			http.Error(w, "Login failed", http.StatusUnauthorized)
			return
		}

		fmt.Fprintf(w, "Welcome, %s!", user)
	})

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
