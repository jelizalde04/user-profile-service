package controllers

import (
    "net/http"
    "encoding/json"
    "view-user-profile-microservice/models"
    "view-user-profile-microservice/services"
    "github.com/gorilla/mux"
)

// ViewUserProfile godoc
// @Summary View a user's profile
// @Description View the profile of a specific user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /view-profile/{id} [get]
func ViewUserProfile(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    userID := vars["id"]

    user, err := services.ViewUserProfileService(userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
}