// Erik PEtrosyan ©
package web

import (
	"net/http"
	"shtem-web/sources/internal/adapters/web/middlewares"
	"shtem-web/sources/internal/core/ports"

	"github.com/gin-gonic/gin"
)

const faviconFile = "svg/logo-32.svg"

func NewWEBRouter(handler ports.WEBHandler) *gin.Engine {

	r := gin.Default()
	middlewares.ApplyCommonMiddlewares(r)

	r.GET("/", handler.Home("home.html"))
	r.GET("/shtems", func(ctx *gin.Context) { ctx.Redirect(http.StatusPermanentRedirect, "/") })
	r.GET("/shtems/:shtemName", handler.Shtems("quiz.html"))
	r.GET("/about", handler.About("about.html"))

	r.StaticFileFS("/favicon.ico", faviconFile, handler.Static())

	st := r.Group("/assets", middlewares.PreventListing(handler.Page404(), "/assets", "/assets/"))
	{
		st.StaticFS("/", handler.Static())
	}

	r.NoRoute(handler.Page404())
	r.NoMethod(func(ctx *gin.Context) { ctx.String(http.StatusMethodNotAllowed, "") })
	return r
}
