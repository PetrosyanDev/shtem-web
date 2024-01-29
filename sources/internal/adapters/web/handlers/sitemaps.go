package handlers

import (
	"log"
	"net/http"

	"shtem-web/sources/internal/adapters/web/dto"

	"github.com/gin-gonic/gin"
)

func (h *webHandler) SiteMapForAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		siteMapData, err := h.shtemsService.GetSiteMap()
		if err != nil {
			log.Printf("webHandler:SiteMap (%v)", err.RawError())
			ctx.XML(http.StatusInternalServerError, nil)
			return
		}
		dto.WriteXMLResponse(ctx, "sitemap.xml", &siteMapData)
	}
}
