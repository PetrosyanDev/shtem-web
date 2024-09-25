package handlers

import (
	"log"
	"shtem-web/sources/internal/adapters/web/dto"

	"github.com/gin-gonic/gin"
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
