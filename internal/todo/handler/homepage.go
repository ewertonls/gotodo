package handler

import (
	"net/http"

	"github.com/ewertonls/gotodo/internal/todo"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func GetHomePage(ctx *gin.Context, db *bun.DB) {
	todos, err := todo.GetTodos(db)
	if err != nil {
		SendResponse(ctx, http.StatusInternalServerError, "internal_server_error.html", []todo.Todo{})
		return
	}
	SendResponse(ctx, http.StatusOK, "home.html", todos)
}
