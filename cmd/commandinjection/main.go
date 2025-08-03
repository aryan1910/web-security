package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/gitlog", func(w http.ResponseWriter, r *http.Request) {
		branch := r.URL.Query().Get("branch")
		fmt.Println(branch)
		if branch == "" {
			http.Error(w, "Missing repo or branch", http.StatusBadRequest)
			return
		}

		// ⚠️ Dangerous: Unsanitized user input passed to shell
		cmd := exec.Command("sh", "-c", fmt.Sprintf("git log %s", branch))
		fmt.Println(cmd)
		out, err := cmd.CombinedOutput()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %v\n%s", err, out), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Git Log:\n%s", out)
	})

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
