// Erik Petrosyan ©
package ports

import "github.com/gin-gonic/gin"

type TemplatesRepo interface {
	SetStatus(ctx *gin.Context, status int) TemplatesRepo
	Render(ctx *gin.Context, name string, data any)
}
