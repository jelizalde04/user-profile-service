package routes

import (
    "github.com/gorilla/mux"
    "update-user-password-microservice/controllers"
)

// RegisterPasswordRoutes registers routes for password-related operations
func RegisterPasswordRoutes(router *mux.Router) {
    router.HandleFunc("/update-password", controllers.UpdateUserPassword).Methods("PUT")
}