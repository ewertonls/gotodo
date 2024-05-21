package handler

import (
	"net/http"
	"time"

	"github.com/ewertonls/gotodo/internal/todo"
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func PostTodo(ctx *gin.Context, db *bun.DB) {
	var todoReq struct {
		Title       string `form:"title"`
		Description string `form:"description"`
	}
	if err := ctx.Bind(&todoReq); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	now := time.Now()
	newTodo := todo.Todo{
		Title:       todoReq.Title,
		Description: todoReq.Description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	createdTodo, err := todo.CreateTodo(db, newTodo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	SendResponse(ctx, http.StatusOK, "TodoComponent", createdTodo)
}
