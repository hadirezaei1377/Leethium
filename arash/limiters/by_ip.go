package limiters

import (
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

var (
	ipLimiterMap = make(map[string]*rate.Limiter)
	mu           sync.Mutex
)

func ByIP(next http.Handler, refillRate float64, tokenBucketSize int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		limiter := getLimiter(ip, refillRate, tokenBucketSize)
		if limiter.Allow() {
			next.ServeHTTP(w, r)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte(`{"error": "too many requests"}`))
		}
	})
}

func getLimiter(ip string, refillRate float64, tokenBucketSize int) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := ipLimiterMap[ip]
	if !exists {
		limiter = rate.NewLimiter(rate.Limit(refillRate), tokenBucketSize)
		ipLimiterMap[ip] = limiter
	}
	return limiter
}
