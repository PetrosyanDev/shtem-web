package handlers

import (
	"log"
	"shtem-web/sources/internal/adapters/web/dto"
	"shtem-web/sources/internal/core/domain"
	"shtem-web/sources/internal/utils"

	"github.com/gin-gonic/gin"
)

func (h *webHandler) CDN() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := ctx.Param("key")
		utils.Strip(&key)
		group := ctx.Param("group")
		utils.Strip(&group)
		if key == "" || group == "" {
			// PAGE UUPS
			ctx.String(404, "404")
			return
		}
		file := domain.File{
			Key:     key,
			OwnerID: group,
		}
		if err := h.filesService.Download(&file); err != nil {
			log.Printf("webHandler:CDN: (%v)", err)
			// PAGE UUPS
			ctx.String(500, "500")
			return
		}
		dto.WriteFileResponse(ctx, &file)
	}
}
