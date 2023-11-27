// Erik Petrosyan ©
package ports

import (
	"net/http"
	"shtem-web/sources/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type WEBService interface {
	Home(ctx *gin.Context, page string, data *domain.Page)
	Page404(ctx *gin.Context, data *domain.Page)
	Static() http.FileSystem
}

type WEBHandler interface {
	Home(page string) gin.HandlerFunc
	Page404() gin.HandlerFunc
	Static() http.FileSystem
}
