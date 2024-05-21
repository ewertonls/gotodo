package handler

import (
	"net/http"

	"github.com/ewertonls/gotodo/internal/todo"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func GetTodos(ctx *gin.Context, db *bun.DB) {
	todos, err := todo.GetTodos(db)
	if err != nil {
		SendResponse(ctx, http.StatusBadRequest, "internal_server_error.html", gin.H{"error": err.Error()})
		return
	}

	SendResponse(ctx, http.StatusOK, "TodoList", todos)
}
