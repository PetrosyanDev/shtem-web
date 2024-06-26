package handlers

import (
	"log"
	"shtem-web/sources/internal/adapters/web/dto"

	"github.com/gin-gonic/gin"
)

func (h *webHandler) Shtems(page string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		shtems, err := h.shtemsService.GetShtems()
		if err != nil {
			log.Printf("Error while geting shtems: %s", err)
		}

		shtemsFromCategory, err := h.categoriesService.GetCategoriesWithShtems()
		if err != nil {
			log.Printf("Error while GetShtemsByCategoryId: %s", err)
		}

		h.webService.Shtems(ctx, page, dto.ShtemsData(shtems, shtemsFromCategory))
	}
}

func (h *webHandler) Category(page string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		categoryName := ctx.Param("categoryName")

		category, err := h.categoriesService.GetCategoryByLinkName(categoryName)
		if err != nil {
			log.Printf("Error while geting shtems: %s", err)
			h.webService.Page404(ctx, dto.NotFoundData())
			return
		}

		shtemarans, err := h.categoriesService.GetShtemsByCategoryLinkName(categoryName)
		if err != nil {
			log.Printf("Error while geting shtems: %s", err)
			h.webService.Page404(ctx, dto.NotFoundData())
			return
		}

		h.webService.Category(ctx, page, dto.CategoryData(category, shtemarans))
	}
}
