// Erik Petrosyan ©
package repositories

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"shtem-web/sources/internal/core/ports"

	"github.com/gin-gonic/gin"
)

const (
	templatesRootDir      = "templates"
	templateComponentsDir = "components"
)

type htmlTemplates struct {
	*template.Template
}

var templateFuncs = map[string]any{}

func (t *htmlTemplates) SetStatus(ctx *gin.Context, status int) ports.TemplatesRepo {
	ctx.Set("rspStatus", status)
	return t
}

func (t *htmlTemplates) Render(ctx *gin.Context, name string, data any) {
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	status := ctx.GetInt("rspStatus")
	if status == 0 {
		ctx.Status(http.StatusOK)
	} else {
		ctx.Status(status)
	}
	if err := t.ExecuteTemplate(ctx.Writer, name, data); err != nil {
		txt := `
			<h1>Template error!</h1>
		`

		ctx.String(http.StatusInternalServerError, txt)
		log.Printf("Render: %v", err)
	}
}

func NewHTMLTemplates(emb *embed.FS) (*htmlTemplates, error) {
	log.Println("loading templates")
	templates, err := fs.Sub(emb, templatesRootDir)
	if err != nil {
		return nil, err
	}
	components, err := fs.Sub(templates, templateComponentsDir)
	if err != nil {
		return nil, err
	}
	t, err := template.New("main").Funcs(templateFuncs).ParseFS(templates, "*.html")
	if err != nil {
		return nil, err
	}
	t, err = t.ParseFS(components, "*.html")
	if err != nil {
		return nil, err
	}
	return &htmlTemplates{t}, nil
}
