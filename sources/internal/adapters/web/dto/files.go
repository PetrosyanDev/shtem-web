package dto

import (
	"bytes"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func WriteXMLResponse(ctx *gin.Context, name string, data *[]byte) {
	ctx.Header("Content-Type", "application/xml; charset=utf-8")
	http.ServeContent(ctx.Writer, ctx.Request, name, time.Now().UTC(), bytes.NewReader(*data))
}
