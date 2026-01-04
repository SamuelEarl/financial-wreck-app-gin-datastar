package routes

import (
	"net/http"

	"financialwreck.com/app/views"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Create a Gin router with default middleware (logger and recovery)
	router := gin.Default()

	// Define a simple GET endpoint
	router.GET("/ping", ping)
	router.GET("/hello", hello)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	return router
}

// This function will render the templ component into a gin context's Response Writer.
func render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

func ping(c *gin.Context) {
	// Return JSON response
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func hello(c *gin.Context) {
	render(c, http.StatusOK, views.Hello("John"))
}
