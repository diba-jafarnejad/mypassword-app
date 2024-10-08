package handlers

import (
	"encoding/json"
	"mypassword-app/backend/models"
	"mypassword-app/backend/utils"
	"net/http"
)

// In-memory storage
var passwordVault = []models.PasswordEntry{}

// CreatePassword creates a new password entry
func CreatePassword(w http.ResponseWriter, r *http.Request) {
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
	json.NewEncoder(w).Encode(passwordVault)
}

// AuthenticateUser is a dummy authentication function
func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
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
