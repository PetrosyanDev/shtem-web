// HRACH_DEV © iMed Cloud Services, Inc.
package handlers

import (
	"net/http"
	"shtem-web/sources/internal/adapters/web/dto"
	"shtem-web/sources/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type webHandler struct {
	webService ports.WEBService
}

func (h *webHandler) Home(page string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h.webService.Home(ctx, page, dto.HomeData())
	}
}

func (h *webHandler) Shtems(page string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h.webService.Home(ctx, page, dto.ShtemsData())
	}

}
func (h *webHandler) About(page string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h.webService.Home(ctx, page, dto.AboutData())
	}
}

func (h *webHandler) Page404() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h.webService.Page404(ctx, dto.NotFoundData())
	}
}

func (h *webHandler) Static() http.FileSystem {
	return h.webService.Static()
}

func NewWEBHandler(webService ports.WEBService) *webHandler {
	return &webHandler{webService}
}
