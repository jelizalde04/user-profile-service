package main

import (
    "log"
    "net/http"
    "os"
    "time"
    "list-user-profiles-microservice/docs"
    "list-user-profiles-microservice/middleware"
    "list-user-profiles-microservice/routes"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    httpSwagger "github.com/swaggo/http-swagger"
)

// @title List User Profiles Microservice API
// @version 1.0
// @description Microservicio para listar todos los perfiles de usuario.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@example.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:6005
// @BasePath /
func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error cargando el archivo .env")
    }

    router := mux.NewRouter()

    // Middlewares
    router.Use(middleware.LoggingMiddleware)
    router.Use(middleware.JWTMiddleware)
    router.Use(middleware.CORSMiddleware)
    router.Use(middleware.RateLimitMiddleware)

    // Rutas
    routes.RegisterUserProfileRoutes(router)

    // Documentación Swagger
    router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

    srv := &http.Server{
        Handler:      router,
        Addr:         ":6005",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    log.Println("Servidor corriendo en el puerto 6005")
    log.Fatal(srv.ListenAndServe())
}
