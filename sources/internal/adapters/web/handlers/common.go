// HRACH_DEV © iMed Cloud Services, Inc.
package handlers

import (
	"log"
	"net/http"
	"shtem-web/sources/internal/adapters/web/dto"
	"shtem-web/sources/internal/core/ports"

	"github.com/gin-gonic/gin"
)

type webHandler struct {
	webService    ports.WEBService
	shtemsService ports.ShtemsService
}

func (h *webHandler) Home(page string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h.webService.Home(ctx, page, dto.HomeData())
	}
}

func (h *webHandler) Shtems(page string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		names, err := h.shtemsService.GetShtemNames()
		if err != nil {
			log.Printf("Error while geting shtems: %s", err)
		}

		log.Println(names)

		h.webService.Shtems(ctx, page, dto.ShtemsData(names))
	}
}

func (h *webHandler) Quiz(page string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		names, err := h.shtemsService.GetShtemNames()
		if err != nil {
			log.Printf("Error while geting shtems: %s", err)
			return
		}

		shtemName := ctx.Param("shtemName")

		for _, val := range names {
			if val.LinkName == shtemName {
				h.webService.Shtems(ctx, page, dto.QuizData())
				return
			}
		}

		h.webService.Page404(ctx, dto.NotFoundData())
	}
}

func (h *webHandler) About(page string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h.webService.About(ctx, page, dto.AboutData())
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

func NewWEBHandler(webService ports.WEBService, shtemsService ports.ShtemsService) *webHandler {
	return &webHandler{webService, shtemsService}
}
