package controllers

import (
    "encoding/json"
    "net/http"
)

// UserProfile representa un perfil de usuario
type UserProfile struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

// GetAllUserProfiles maneja la solicitud para obtener todos los perfiles de usuario
func GetAllUserProfiles(w http.ResponseWriter, r *http.Request) {
    profiles := []UserProfile{
        {ID: "1", Name: "Juan Pérez", Email: "juan@example.com"},
        {ID: "2", Name: "Ana Gómez", Email: "ana@example.com"},
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(profiles)
}
