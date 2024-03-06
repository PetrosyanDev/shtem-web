package dto

import (
	"bytes"
	"fmt"
	"net/http"
	"shtem-web/sources/internal/core/domain"
	"time"

	"github.com/gin-gonic/gin"
)

func WriteXMLResponse(ctx *gin.Context, name string, data *[]byte) {
	ctx.Header("Content-Type", "application/xml; charset=utf-8")
	http.ServeContent(ctx.Writer, ctx.Request, name, time.Now().UTC(), bytes.NewReader(*data))
}

func WriteFileResponse(ctx *gin.Context, file *domain.File) {
	ctx.Header("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", file.Key))
	http.ServeContent(ctx.Writer, ctx.Request, file.Key, file.CreatedAt, bytes.NewReader(file.Bytes()))
}
