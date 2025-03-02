package routes

import (
    "github.com/gorilla/mux"
    "update-user-email-microservice/controllers"
)

// RegisterEmailRoutes registers routes for email-related operations
func RegisterEmailRoutes(router *mux.Router) {
    router.HandleFunc("/update-email", controllers.UpdateUserEmail).Methods("PUT")
}