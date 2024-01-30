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

		// categories, err := h.categoriesService.GetCategories()
		// if err != nil {
		// 	log.Printf("Error while GetCategories: %s", err)
		// }

		shtemsFromCategory, err := h.categoriesService.GetCategoriesWithShtems()
		if err != nil {
			log.Printf("Error while GetShtemsByCategoryId: %s", err)
		}

		h.webService.Shtems(ctx, page, dto.ShtemsData(shtems, shtemsFromCategory))
	}
}
