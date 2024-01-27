// Erik Petrosyan ©
package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

func PreventListing(fx func(ctx *gin.Context), dirs ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, dir := range dirs {
			log.Println(c.Request.RequestURI == dir)
			if c.Request.RequestURI == dir {
				fx(c)
				c.Abort()
				break
			}
		}
	}
}
