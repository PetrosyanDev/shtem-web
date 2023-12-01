// Erik Petrosyan ©
package services

import (
	"embed"
	"io/fs"
	"net/http"
	"shtem-web/sources/internal/core/domain"
	"shtem-web/sources/internal/core/ports"

	"github.com/gin-gonic/gin"
)

const staticDir = "assets"

type webService struct {
	assets    http.FileSystem
	templates ports.TemplatesRepo
}

func (s *webService) Home(ctx *gin.Context, page string, data *domain.Page) {
	s.templates.Render(ctx, page, data)
}
func (s *webService) Shtems(ctx *gin.Context, page string, data *domain.Page) {
	s.templates.Render(ctx, page, data)
}
func (s *webService) About(ctx *gin.Context, page string, data *domain.Page) {
	s.templates.Render(ctx, page, data)
}

func (s *webService) Page404(ctx *gin.Context, data *domain.Page) {
	s.templates.SetStatus(ctx, http.StatusNotFound).Render(ctx, "404.html", data)
}

func (s *webService) Static() http.FileSystem {
	return s.assets
}

func NewWEBService(emb *embed.FS, template ports.TemplatesRepo) (*webService, error) {
	assets, err := fs.Sub(emb, staticDir)
	if err != nil {
		return nil, err
	}
	return &webService{http.FS(assets), template}, nil
}
