package limiters

import (
	"net/http"

	"golang.org/x/time/rate"
)

func ByAppKey(next http.Handler, refillRate float64, tokenBucketSize int) http.Handler {
	limiter := rate.NewLimiter(rate.Limit(refillRate), tokenBucketSize)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() {
			next.ServeHTTP(w, r)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte(`{"error": "too many requests"}`))
		}
	})
}
