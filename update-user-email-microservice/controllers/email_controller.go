package controllers

import (
    "net/http"
    "encoding/json"
    "update-user-email-microservice/models"
    "update-user-email-microservice/services"
)

// UpdateUserEmail godoc
// @Summary Update a user's email
// @Description Update the email of a specific user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User object"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /update-email [put]
func UpdateUserEmail(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := services.UpdateUserEmailService(user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Email updated successfully"})
}