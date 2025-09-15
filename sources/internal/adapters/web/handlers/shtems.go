package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"
	"shtem-web/sources/internal/adapters/web/dto"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	mathSponsor = "https://mathmind.am/index/trackVisit/shtemaran"
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
	// tracking
	clientID, _ := ctx.Cookie("cid")
	ipH := hashIP(clientIP(ctx))
	ua := ctx.GetHeader("User-Agent")

	err := h.sponsorHitsService.InsertSponsorHit(path, sponsorURL, clientID, ipH, ua)
	if err != nil {
		log.Println(err)
		return
	}

	// redirect
	ctx.Redirect(302, sponsorURL)
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

func clientIP(c *gin.Context) string {
	if ip := c.GetHeader("CF-Connecting-IP"); ip != "" {
		return ip
	}
	if xff := c.GetHeader("X-Forwarded-For"); xff != "" {
		return strings.TrimSpace(strings.Split(xff, ",")[0])
	}
	return c.ClientIP()
}

func hashIP(ip string) string {
	salt := os.Getenv("IP_HASH_SALT")
	sum := sha256.Sum256([]byte(salt + "|" + ip))
	return hex.EncodeToString(sum[:])
}
