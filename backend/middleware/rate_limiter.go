package middleware

import (
	"net/http"
	"shopping-cart/backend/models"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	visitors = make(map[string]*visitor)
	mu       sync.Mutex
)

// cleanup removes old entries from visitors map
func cleanup() {
	mu.Lock()
	for ip, v := range visitors {
		if time.Since(v.lastSeen) > 3*time.Hour {
			delete(visitors, ip)
		}
	}
	mu.Unlock()
}

// RateLimiter creates a new rate limiter middleware
// limit: number of requests
// burst: number of requests that can exceed the limit
func RateLimiter(limit float64, burst int) gin.HandlerFunc {
	go func() {
		for {
			time.Sleep(time.Hour)
			cleanup()
		}
	}()

	return func(c *gin.Context) {
		ip := c.ClientIP()
		mu.Lock()
		v, exists := visitors[ip]
		if !exists {
			v = &visitor{
				limiter: rate.NewLimiter(rate.Limit(limit), burst),
			}
			visitors[ip] = v
		}
		v.lastSeen = time.Now()
		mu.Unlock()

		if !v.limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, models.APIResponse{
				Success: false,
				Error: &models.ErrorResponse{
					Code:    models.ErrCodeRateLimitExceeded,
					Message: "Rate limit exceeded",
					Details: "Too many requests, please try again later",
				},
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
