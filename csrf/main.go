package main

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func main() {
	r := gin.Default()

	// Setup session store
	store := cookie.NewStore([]byte("super-secret-key"))
	r.Use(sessions.Sessions("mysession", store))

	// CSRF middleware allows the GET method and validates the csrf token for post request
	r.Use(csrf.Middleware(csrf.Options{
		Secret: "a-32-byte-secret-key-for-production!",
		ErrorFunc: func(c *gin.Context) {
			c.String(http.StatusForbidden, "CSRF token mismatch")
			c.Abort()
		},
	}))

	// Load HTML template
	tmplPath := filepath.Join("templates", "form.html")
	tmpl := template.Must(template.ParseFiles(tmplPath))
	r.SetHTMLTemplate(tmpl)

	// Routes
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.html", gin.H{
			"csrfToken": csrf.GetToken(c),
		})
	})

	r.POST("/submit", func(c *gin.Context) {
		name := c.PostForm("name")
		c.String(http.StatusOK, "Hello, %s! Your form is submitted securely.", name)
	})

	r.Run(":8080")
}
