package handlers

import (
	"encoding/json"
	"mypassword-app/backend/models"
	"mypassword-app/backend/utils"
	"net/http"
)

// In-memory storage
var passwordVault = []models.PasswordEntry{}

// AddCORSHeaders adds the necessary CORS headers to the response
func AddCORSHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                            // Allow all origins
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")          // Allow specific methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Allow specific headers
}

// Handle preflight OPTIONS requests globally
func Preflight(w http.ResponseWriter, r *http.Request) {
	// Add CORS headers to the preflight response
	AddCORSHeaders(w)
	w.WriteHeader(http.StatusOK) // Respond with OK status for OPTIONS requests
}

// CreatePassword creates a new password entry
func CreatePassword(w http.ResponseWriter, r *http.Request) {
	// Add CORS headers for all requests
	AddCORSHeaders(w)

	// Handle preflight OPTIONS request
	if r.Method == http.MethodOptions {
		Preflight(w, r)
		return
	}

	// Main logic for POST request
	var entry models.PasswordEntry
	_ = json.NewDecoder(r.Body).Decode(&entry)

	encryptedPassword, err := utils.Encrypt(entry.Password)
	if err != nil {
		http.Error(w, "Unable to encrypt password", http.StatusInternalServerError)
		return
	}
	entry.Password = encryptedPassword

	passwordVault = append(passwordVault, entry)
	json.NewEncoder(w).Encode(entry)
}

// GetPasswords returns all passwords
func GetPasswords(w http.ResponseWriter, r *http.Request) {
	// Add CORS headers for all requests
	AddCORSHeaders(w)

	// Handle preflight OPTIONS request
	if r.Method == http.MethodOptions {
		Preflight(w, r)
		return
	}

	// Main logic for GET request
	json.NewEncoder(w).Encode(passwordVault)
}

// AuthenticateUser is a dummy authentication function
func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	// Add CORS headers for all requests
	AddCORSHeaders(w)

	// Handle preflight OPTIONS request
	if r.Method == http.MethodOptions {
		Preflight(w, r)
		return
	}

	// Main logic for POST request (authentication)
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	// For simplicity, we're just checking hardcoded credentials
	if user.Username == "admin" && user.Password == "password" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Authenticated"))
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}
