package routes

import (
    "list-user-profiles-microservice/controllers"
    "github.com/gorilla/mux"
)

// RegisterUserProfileRoutes registra las rutas para obtener los perfiles de usuario
func RegisterUserProfileRoutes(router *mux.Router) {
    router.HandleFunc("/users", controllers.GetAllUserProfiles).Methods("GET")
}
