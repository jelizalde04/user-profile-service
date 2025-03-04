package middleware

import (
	"log"
	"net/http"
	"os"
	"sync"
	"time"
	"encoding/json"
)

// RateLimiter estructura para limitar el número de solicitudes por IP
type RateLimiter struct {
	visitors map[string]*Visitor
	mu       sync.Mutex
	limit    int
	window   time.Duration
}

// Visitor almacena la información de cada IP
type Visitor struct {
	lastSeen time.Time
	count    int
}

// NewRateLimiter crea un nuevo RateLimiter
func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	rl := &RateLimiter{
		visitors: make(map[string]*Visitor),
		limit:    limit,
		window:   window,
	}
	go rl.cleanup()
	return rl
}

// Middleware para aplicar Rate Limiting
func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		rl.mu.Lock()
		v, exists := rl.visitors[ip]
		if !exists {
			v = &Visitor{lastSeen: time.Now(), count: 1}
			rl.visitors[ip] = v
		} else {
			if time.Since(v.lastSeen) > rl.window {
				v.count = 1
				v.lastSeen = time.Now()
			} else {
				v.count++
			}
		}
		rl.mu.Unlock()

		if v.count > rl.limit {
			logSuspiciousIP(ip)
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Limpieza de visitantes inactivos
func (rl *RateLimiter) cleanup() {
	for {
		time.Sleep(rl.window)
		rl.mu.Lock()
		for ip, v := range rl.visitors {
			if time.Since(v.lastSeen) > rl.window {
				delete(rl.visitors, ip)
			}
		}
		rl.mu.Unlock()
	}
}

// Función para registrar IPs sospechosas en un archivo
func logSuspiciousIP(ip string) {
	file, err := os.OpenFile("ddos_attempts.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	entry := map[string]string{"ip": ip, "timestamp": time.Now().Format(time.RFC3339)}
	jsonEntry, _ := json.Marshal(entry)
	file.WriteString(string(jsonEntry) + "\n")
}
