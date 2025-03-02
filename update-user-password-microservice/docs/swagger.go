package docs

import (
	"github.com/swaggo/swag"
)

// @title Update User Email API
// @version 1.0
// @description Microservice for updating user email.
// @host localhost:6002
// @BasePath /
// @schemes http

func SwaggerInfo() *swag.Spec {
	return &swag.Spec{
		Version:          "1.0",
		Title:            "Update User Email API",
		Description:      "Microservice for updating user email.",
		Host:             "localhost:6002",
		BasePath:         "/",
		Schemes:          []string{"http"},
	}
}
