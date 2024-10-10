package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	ratelimiter "github.com/khaaleoo/gin-rate-limiter/core"
)

var rateLimiterOption = ratelimiter.RateLimiterOption{
	Limit: 0.1,
	Burst: 5,
	Len:   1 * time.Minute,
}

func RateLimiterMiddleware() gin.HandlerFunc {
	return ratelimiter.RequireRateLimiter(ratelimiter.RateLimiter{
		RateLimiterType: ratelimiter.IPRateLimiter,
		Key:             "ip_rate_limiter",
		Option:          rateLimiterOption,
	})
}
