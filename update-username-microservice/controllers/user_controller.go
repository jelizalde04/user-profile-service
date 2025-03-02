package controllers

import (
    "net/http"
    "encoding/json"
    "update-username-microservice/models"
    "update-username-microservice/services"
)

// UpdateUsername godoc
// @Summary Update a user's username
// @Description Update the username of a specific user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User object"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /update-username [put]
func UpdateUsername(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := services.UpdateUsernameService(user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Username updated successfully"})
}