package middleware

import (
    "log"
    "net/http"
    "github.com/golang-jwt/jwt/v4"
    "strings"
    "golang.org/x/time/rate"
)

// LoggingMiddleware logs requests
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

// JWTMiddleware 
func JWTMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }

        tokenString := strings.Split(authHeader, " ")[1]
        _, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte("mysecretkey"), nil
        })

        if err != nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    })
}

// RateLimitMiddleware DDOS y EDOS
var limiter = rate.NewLimiter(1, 3)

func RateLimitMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !limiter.Allow() {
            http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
            return
        }
        next.ServeHTTP(w, r)
    })
}
