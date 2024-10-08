package router

import (
	"mypassword-app/backend/handlers"

	"github.com/gorilla/mux"
)

// SetupRouter initializes the API routes
func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	// Authentication route
	r.HandleFunc("/auth", handlers.AuthenticateUser).Methods("POST")

	// Password routes
	r.HandleFunc("/passwords", handlers.CreatePassword).Methods("POST")
	r.HandleFunc("/passwords", handlers.GetPasswords).Methods("GET")

	return r
}
