package handlers

import (
	"encoding/json"
	"net/http"
	"user-service/internal/models"
)

// GetUserByID returns a single user by their ID.
// For now, it returns a hardcoded user.
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	// In a real application, you would fetch the user from a database
	// based on the ID from the request URL.
	user := models.User{
		ID:    "1",
		PayID: "user1@superapp",
		Name:  "Test User",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}