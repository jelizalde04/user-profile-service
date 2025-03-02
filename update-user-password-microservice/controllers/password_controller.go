package controllers

import (
    "net/http"
    "encoding/json"
    "update-user-password-microservice/models"
    "update-user-password-microservice/services"
)

// UpdateUserPassword godoc
// @Summary Update a user's password
// @Description Update the password of a specific user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User object"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /update-password [put]
func UpdateUserPassword(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := services.UpdateUserPasswordService(user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Password updated successfully"})
}