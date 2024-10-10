package main

import (
	"log"
	"mypassword-app/backend/middleware"
	"mypassword-app/backend/router"
	"net/http"
)

func main() {
	r := router.SetupRouter() // Initialize the router
	handlerWithCORS := middleware.CorsMiddleware(r)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", handlerWithCORS) // Start the server
}
