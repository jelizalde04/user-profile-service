package routes

import (
    "github.com/gorilla/mux"
    "update-username-microservice/controllers"
)

// RegisterUserRoutes registers routes for user-related operations
func RegisterUserRoutes(router *mux.Router) {
    router.HandleFunc("/update-username", controllers.UpdateUsername).Methods("PUT")
}