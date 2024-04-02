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
	webService        ports.WEBService
	questionsService  ports.QuestionsService
	shtemsService     ports.ShtemsService
	categoriesService ports.CategoriesService
	filesService      ports.FilesService
}

func (h *webHandler) Home(page string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		shtems, err := h.shtemsService.GetShtems()
		if err != nil {
			log.Printf("Error while geting shtems: %s", err)
		}

		shtemsFromCategory, err := h.categoriesService.GetCategoriesWithShtems()
		if err != nil {
			log.Printf("Error while GetShtemsByCategoryId: %s", err)
		}

		// h.webService.Shtems(ctx, page, dto.ShtemsData(shtems, shtemsFromCategory))

		h.webService.Home(ctx, page, dto.HomeData(shtems, shtemsFromCategory))
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

func (h *webHandler) StaticUploads() http.FileSystem {
	return h.webService.StaticUploads()
}

func NewWEBHandler(
	webService ports.WEBService,
	questionsService ports.QuestionsService,
	shtemsService ports.ShtemsService,
	categoriesService ports.CategoriesService,
	filesService ports.FilesService,
) *webHandler {
	return &webHandler{webService, questionsService, shtemsService, categoriesService, filesService}
}
