package middlewares

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/gin-gonic/gin"
)

const clientCookie = "cid"

func EnsureClientID() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := c.Cookie(clientCookie); err != nil {
			var b [16]byte
			_, _ = rand.Read(b[:])
			val := hex.EncodeToString(b[:])
			c.SetCookie(clientCookie, val, int((365 * 24 * time.Hour).Seconds()), "/", "", true, true)
		}
		c.Next()
	}
}
