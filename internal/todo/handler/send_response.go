package handler

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func SendResponse(ctx *gin.Context, status int, templateName string, data any) {
	accept := ctx.GetHeader("Accept")
	if strings.Contains(accept, "application/json") {
		ctx.JSON(status, data)
	} else {
		ctx.HTML(status, templateName, data)
	}
}
