package routes

import (
	"net/http"

	"financialwreck.com/app/views"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	datastar "github.com/starfederation/datastar/sdk/go"
)

func SetupRouter() *gin.Engine {
	// Create a Gin router with default middleware (logger and recovery)
	router := gin.Default()

	// Serve physical files from the /static folder
	router.Static("/static", "./static")

	// Serve virtual CSS from Templ's built-in "css" blocks.
	// Create a middleware that gathers all the Templ CSS (from the `css` functions in your templates).
	// Wrap the Templ CSS Middleware so Gin can use it.
	// This will serve all your "css" blocks at /templ/style.css
	router.GET("/templ/style.css", gin.WrapH(templ.NewCSSMiddleware(nil)))

	// Define a simple GET endpoint
	router.GET("/", home)
	router.GET("/ping", ping)
	router.GET("/hello", hello)
	router.GET("/counter", counter)
	router.POST("/increment", increment)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	return router
}

func home(c *gin.Context) {
	views.RenderPage(c, 200, "Home Page", views.Home())
}

func ping(c *gin.Context) {
	// Return JSON response
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func hello(c *gin.Context) {
	views.RenderPage(c, http.StatusOK, "Hello", views.Hello("John"))
}

func counter(c *gin.Context) {
	views.RenderPage(c, 200, "Counter", views.Counter(0))
}

// Create a struct that matches the keys you defined in your data-signals attribute in the counter.templ component.
type CounterSignals struct {
	Count int `json:"count"`
}

func increment(c *gin.Context) {
	// 1. Parse the incoming signals from the frontend
	var signals CounterSignals
	err := c.ShouldBindJSON(&signals)
	if err != nil {
		c.Status(400)
		return
	}

	// 2. Perform your logic
	signals.Count++

	// 3. Setup the Datastar Server-Sent Event (SSE) stream
	sse := datastar.NewSSE(c.Writer, c.Request)

	// 4. Render the Templ component to a string and send it as a fragment. (i.e. Push the updated fragment back to the browser.)
	// This replaces the HTML and updates the frontend signals. (i.e. Datastar will look for the ID "container" and replace it.)
	sse.MergeFragmentTempl(views.Counter(signals.Count))

	// // 4. (Alternative) Advanced: Sending "Only" a Signal Update
	// // Sometimes you don't want to re-render the HTML; you just want to update a value in the browser's memory (the signal). Datastar allows you to send a signal update without a fragment using the MergeSignals() method.

	// // Marshal the struct directly
	// payload, _ := json.Marshal(signals)

	// // This updates the 'count' variable in the browser
	// // without changing any HTML elements!
	// sse.MergeSignals(payload)
}
