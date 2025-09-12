// Erik Petrosyan ©
package ports

import (
	"net/http"
	"shtem-web/sources/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type WEBService interface {
	Home(ctx *gin.Context, page string, data *domain.Page)
	Blog(ctx *gin.Context, page string, data *domain.Page)
	Shtems(ctx *gin.Context, page string, data *domain.Page)
	About(ctx *gin.Context, page string, data *domain.Page)
	Quiz(ctx *gin.Context, page string, data *domain.Page)
	SingleShtem(ctx *gin.Context, page string, data *domain.Page)
	Category(ctx *gin.Context, page string, data *domain.Page)
	Page404(ctx *gin.Context, data *domain.Page)
	Static() http.FileSystem
	StaticUploads() http.FileSystem
}

type WEBHandler interface {
	Home(page string) gin.HandlerFunc
	Blog(page string) gin.HandlerFunc
	Shtems(page string) gin.HandlerFunc
	Quiz(page string) gin.HandlerFunc
	SingleShtem(page string) gin.HandlerFunc
	SingleShtemSponsor() gin.HandlerFunc
	Category(page string) gin.HandlerFunc
	About(page string) gin.HandlerFunc
	EmailSubmit() gin.HandlerFunc
	Page404() gin.HandlerFunc
	Static() http.FileSystem
	CDN() gin.HandlerFunc
	StaticUploads() http.FileSystem
	SiteMapForAll() gin.HandlerFunc
}
