package handler

import (
	"net/http"
	"strconv"

	"github.com/ewertonls/gotodo/internal/todo"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func PutTodo(ctx *gin.Context, db *bun.DB) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		SendResponse(ctx, http.StatusBadRequest, "internal_server_error.html", err.Error())
	}

	completedStr := ctx.Param("completed")
	completed := false
	if completedStr == "true" {
		completed = true
	}

	todoUpdate := todo.Todo{ID: uint(id)}
	if completed {
		todoUpdate.MarkAsCompleted()
	} else {
		todoUpdate.MarkAsUncompleted()
	}

	updatedTodo, err := todo.UpdateTodo(db, todoUpdate)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "internal_server_error.html", err.Error())
		return
	}

	SendResponse(ctx, http.StatusOK, "TodoComponent", updatedTodo)
}
