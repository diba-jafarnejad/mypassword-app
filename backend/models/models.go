package models

// PasswordEntry represents a stored password entry
type PasswordEntry struct {
	ID       int    `json:"id"`
	Account  string `json:"account"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// User represents a simple user authentication model
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
