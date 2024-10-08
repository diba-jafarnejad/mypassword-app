package main

import (
	"log"
	"mypassword-app/backend/router"
	"net/http"
)

func main() {
	r := router.SetupRouter() // Initialize the router
	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r) // Start the server
}
