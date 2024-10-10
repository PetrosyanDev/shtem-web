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
	emailsService     ports.EmailsService
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

func (h *webHandler) Blog(page string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h.webService.Blog(ctx, page, dto.BlogData())
	}
}
func (h *webHandler) About(page string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h.webService.About(ctx, page, dto.AboutData())
	}
}
func (h *webHandler) EmailSubmit() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req struct {
			Email string `json:"email" binding:"required,email"`
		}

		// Bind the JSON request to req
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "Invalid JSON format"})
			return
		}

		// Attempt to insert the email and handle errors
		err := h.emailsService.InsertEmail(req.Email)
		if err != nil {
			if err.GetMessage() == "Email already exists" {
				ctx.JSON(http.StatusConflict, gin.H{"success": false, "error": "This email is already registered."})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "Something went wrong! Please try again later."})
			}
			return
		}

		// Email submitted successfully
		ctx.JSON(http.StatusOK, gin.H{"success": true, "message": "Email submitted successfully"})
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
	emailsService ports.EmailsService,
	filesService ports.FilesService,
) *webHandler {
	return &webHandler{webService, questionsService, shtemsService, categoriesService, emailsService, filesService}
}
