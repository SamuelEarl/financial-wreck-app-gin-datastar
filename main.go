package main

import "financialwreck.com/app/routes"

func main() {
	router := routes.SetupRouter()

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	router.Run()
}
