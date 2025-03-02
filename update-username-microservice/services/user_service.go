package services

import (
    "update-username-microservice/models"
    "update-username-microservice/utils"
)

// UpdateUsernameService updates a user's username
func UpdateUsernameService(user models.User) error {
    // Encrypt the username
    encryptedUsername, err := utils.Encrypt(user.Username)
    if err != nil {
        return err
    }

    // Logic to update the username in the database
    // ...

    return nil
}