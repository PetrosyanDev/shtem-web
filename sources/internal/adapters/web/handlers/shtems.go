package handlers

import (
	"log"
	"shtem-web/sources/internal/adapters/web/dto"

	"github.com/gin-gonic/gin"
)

const (
	mathSponsor = "https://mathmind.am/?utm_source=shtemaran.am"
)

func (h *webHandler) SingleShtem(page string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		shtemName := ctx.Param("shtemName")

		shtemaran, err := h.shtemsService.GetShtemByLinkName(shtemName)
		if err != nil {
			log.Printf("Error while geting shtems: %s", err)
			h.webService.Page404(ctx, dto.NotFoundData())
			return
		}

		category, err := h.categoriesService.GetCategoryByShtemLinkName(shtemName)
		if err != nil {
			log.Printf("Error while geting category: %s", err)
			h.webService.Page404(ctx, dto.NotFoundData())
			return
		}

		h.webService.SingleShtem(ctx, page, dto.SingleShtemData(category, shtemaran))
	}
}

func (h *webHandler) SingleShtemSponsor() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		shtemName := ctx.Param("shtemName")

		_, err := h.shtemsService.GetShtemByLinkName(shtemName)
		if err != nil {
			h.webService.Page404(ctx, dto.NotFoundData())
			return
		}

		switch shtemName {
		case "matem-1":
			h.sponsor(ctx, shtemName, mathSponsor)
			return
		case "matem-2":
			h.sponsor(ctx, shtemName, mathSponsor)
			return
		default:
			ctx.Redirect(301, "/")
			return
		}
	}
}

func (h *webHandler) sponsor(ctx *gin.Context, path string, sponsorURL string) {
	h.tgClient.NotifyOnSponsor(path, sponsorURL)
	ctx.Redirect(301, sponsorURL)
}

func (h *webHandler) Quiz(page string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		shtemName := ctx.Param("shtemName")

		_, err := h.shtemsService.GetShtemByLinkName(shtemName)
		if err != nil {
			log.Printf("Error while geting shtems: %s", err)
			h.webService.Page404(ctx, dto.NotFoundData())
			return
		}

		h.webService.Shtems(ctx, page, dto.QuizData())
	}
}
