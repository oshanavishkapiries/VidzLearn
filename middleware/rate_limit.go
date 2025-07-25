package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"
)

type rateLimiter struct {
	mu        sync.Mutex
	tokens    float64
	lastCheck time.Time
}

const (
	rateLimit     = 5  // tokens per second
	burstCapacity = 10 // max burst
	cleanupWindow = 5 * time.Minute
)

var (
	limiters   = make(map[string]*rateLimiter)
	limitersMu sync.Mutex
)

// RateLimitMiddleware limits requests per IP
func RateLimitMiddleware(next http.Handler) http.Handler {
	go cleanupStaleLimiters()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "Rate limiter failed to parse IP", http.StatusInternalServerError)
			return
		}

		limiter := getLimiter(ip)
		if !limiter.allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func getLimiter(ip string) *rateLimiter {
	limitersMu.Lock()
	defer limitersMu.Unlock()

	if limiter, exists := limiters[ip]; exists {
		return limiter
	}

	limiter := &rateLimiter{
		tokens:    burstCapacity,
		lastCheck: time.Now(),
	}
	limiters[ip] = limiter
	return limiter
}

func (rl *rateLimiter) allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastCheck).Seconds()
	rl.tokens += elapsed * rateLimit
	if rl.tokens > burstCapacity {
		rl.tokens = burstCapacity
	}
	rl.lastCheck = now

	if rl.tokens >= 1 {
		rl.tokens--
		return true
	}
	return false
}

func cleanupStaleLimiters() {
	for {
		time.Sleep(cleanupWindow)
		limitersMu.Lock()
		now := time.Now()
		for ip, limiter := range limiters {
			if now.Sub(limiter.lastCheck) > cleanupWindow {
				delete(limiters, ip)
			}
		}
		limitersMu.Unlock()
	}
}
