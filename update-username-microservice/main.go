package main

import (
    "log"
    "net/http"
    "os"
    "time"
    "update-username-microservice/docs" // Automatically generated by Swag
    "update-username-microservice/middleware"
    "update-username-microservice/routes"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    httpSwagger "github.com/swaggo/http-swagger"
)

// @title Update Username Microservice API
// @version 1.0
// @description This is a microservice to update a user's username.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@example.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:6001
// @BasePath /
func main() {
    // Load environment variables
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file")
    }

    // Configure the router
    router := mux.NewRouter()

    // Middleware
    router.Use(middleware.LoggingMiddleware)
    router.Use(middleware.CORSMiddleware)
    router.Use(middleware.JWTAuthMiddleware)

    // Rate Limiting (100 requests per 15 minutes)
    rateLimiter := middleware.NewRateLimiter(100, 15*time.Minute)
    router.Use(rateLimiter.Middleware)

    // Routes
    routes.RegisterUserRoutes(router)

    // Swagger UI
    router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

    // Start the server
    port := os.Getenv("PORT")
    if port == "" {
        port = "6001"
    }

    log.Printf("Server started on port %s", port)
    log.Fatal(http.ListenAndServe(":"+port, router))
}
