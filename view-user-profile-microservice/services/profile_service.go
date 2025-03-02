package services

import (
    "view-user-profile-microservice/models"
    "encoding/json"
    "net/http"
)

// ViewUserProfileService retrieves a user's profile
func ViewUserProfileService(userID string) (*models.User, error) {
    // Make a REST call to another microservice (e.g., update-username-microservice)
    resp, err := http.Get("http://update-username-microservice:6001/user/" + userID)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var user models.User
    if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
        return nil, err
    }

    return &user, nil
}