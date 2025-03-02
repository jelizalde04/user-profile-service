package models

// User represents a user entity
type User struct {
    ID       string `json:"id"`
    Password string `json:"password"`
}