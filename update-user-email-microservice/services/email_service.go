package services

import (
    "update-user-email-microservice/models"
    "update-user-email-microservice/utils"
)

// UpdateUserEmailService updates a user's email
func UpdateUserEmailService(user models.User) error {
    // Encrypt the email
    encryptedEmail, err := utils.Encrypt(user.Email)
    if err != nil {
        return err
    }

    // Logic to update the email in the database
    // ...

    return nil
}