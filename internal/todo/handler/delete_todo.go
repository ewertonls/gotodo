package handler

import (
	"net/http"
	"strconv"

	"github.com/ewertonls/gotodo/internal/todo"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func DeleteTodo(ctx *gin.Context, db *bun.DB) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		SendResponse(ctx, http.StatusBadRequest, "internal_server_error.html", err.Error())
		return
	}

	if err := todo.DeleteTodo(db, uint(id)); err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "internal_server_error.html", err.Error())
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}
