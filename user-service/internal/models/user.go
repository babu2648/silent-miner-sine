package models

// User represents a user in the system.
type User struct {
	ID    string `json:"id"`
	PayID string `json:"payId"`
	Name  string `json:"name"`
}