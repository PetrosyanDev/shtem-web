// Erik PEtrosyan ©
package web

import (
	"net/http"
	"shtem-web/sources/internal/adapters/web/middlewares"
	"shtem-web/sources/internal/core/ports"

	"github.com/gin-gonic/gin"
)

const faviconFile = "img/logo-32.png"

func NewWEBRouter(handler ports.WEBHandler) *gin.Engine {

	r := gin.Default()
	middlewares.ApplyCommonMiddlewares(r)

	r.GET("/", handler.Home("quiz.html"))

	r.StaticFileFS("/favicon.ico", faviconFile, handler.Static())

	st := r.Group("/assets", middlewares.PreventListing(handler.Page404(), "/assets", "/assets/"))
	{
		st.StaticFS("/", handler.Static())
	}

	r.NoRoute(handler.Page404())
	r.NoMethod(func(ctx *gin.Context) { ctx.String(http.StatusMethodNotAllowed, "") })
	return r
}
