// Erik Petrosyan ©
package middlewares

import (
	"time"

	connLimit "github.com/aviddiviner/gin-limit"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

const maxConcurrentCon int = 10

var (
	allowedHeaders = []string{
		"X-Forwarded-For", "X-Imedcs-Api-Key", "X-Imedcs-Access-Token", "X-Locale", "X-Timezone",
	}
	allowedMethods = []string{
		"GET", "HEAD", "OPTIONS",
	}

	nonCompressables = []string{
		".png", ".gif", ".jpeg", ".jpg", ".pdf", ".mp4", ".avi", ".mov", ".webp", ".mp3",
	}
)

func ApplyCommonMiddlewares(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: allowedMethods,
		AllowHeaders: allowedHeaders,
		MaxAge:       1 * time.Hour,
	}))
	r.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedExtensions(nonCompressables)))
	r.Use(helmet.DNSPrefetchControl())
	r.Use(helmet.FrameGuard())
	r.Use(helmet.IENoOpen())
	r.Use(helmet.SetHSTS(true, 16768060))
	r.Use(helmet.NoSniff())
	r.Use(helmet.XSSFilter())
	r.Use(helmet.NoCache())
	r.Use(connLimit.MaxAllowed(maxConcurrentCon))
}
