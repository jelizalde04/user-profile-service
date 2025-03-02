package middleware

import (
    "net/http"
    "sync"
    "time"
)

// RateLimiter limits the number of requests from a single IP address
type RateLimiter struct {
    requests map[string]int
    mutex    sync.Mutex
    limit    int
    interval time.Duration
}

// NewRateLimiter creates a new RateLimiter instance
func NewRateLimiter(limit int, interval time.Duration) *RateLimiter {
    return &RateLimiter{
        requests: make(map[string]int),
        limit:    limit,
        interval: interval,
    }
}

// Middleware enforces rate limiting
func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ip := r.RemoteAddr

        rl.mutex.Lock()
        count, exists := rl.requests[ip]
        if !exists {
            rl.requests[ip] = 1
            go rl.resetCounter(ip)
        } else if count >= rl.limit {
            rl.mutex.Unlock()
            http.Error(w, "Too many requests", http.StatusTooManyRequests)
            return
        } else {
            rl.requests[ip]++
        }
        rl.mutex.Unlock()

        next.ServeHTTP(w, r)
    })
}

// resetCounter resets the request count for an IP address
func (rl *RateLimiter) resetCounter(ip string) {
    time.Sleep(rl.interval)
    rl.mutex.Lock()
    delete(rl.requests, ip)
    rl.mutex.Unlock()
}