package docs

import (
	"github.com/swaggo/swag"
)

// @title Update Username API
// @version 1.0
// @description Microservicio para actualizar el nombre de usuario.
// @host localhost:6001
// @BasePath /
// @schemes http

func SwaggerInfo() *swag.Spec {
	return &swag.Spec{
		Version:          "1.0",
		Title:            "Update Username API",
		Description:      "Microservicio para actualizar el nombre de usuario.",
		Host:             "localhost:6001",
		BasePath:         "/",
		Schemes:          []string{"http"},
	}
}
