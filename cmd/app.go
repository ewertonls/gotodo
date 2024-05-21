package main

import (
	"log"

	"github.com/ewertonls/gotodo/internal/data"
	"github.com/ewertonls/gotodo/internal/todo/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := data.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := data.CreateTables(db); err != nil {
		log.Println(err)
	}

	r := gin.Default()

	// Assuming server will run from cmd/ dir
	r.LoadHTMLGlob("../templates/*")

	// Web route
	r.GET("/", func(ctx *gin.Context) {
		handler.GetHomePage(ctx, db)
	})

	// API routes
	api := r.Group("api/")
	api.GET("/todos", func(ctx *gin.Context) {
		handler.GetTodos(ctx, db)
	})

	api.POST("/todos", func(ctx *gin.Context) {
		handler.PostTodo(ctx, db)
	})

	api.PUT("/todos/toggleCompleted/:id/:completed", func(ctx *gin.Context) {
		handler.PutTodo(ctx, db)
	})

	api.DELETE("/todos/:id", func(ctx *gin.Context) {
		handler.DeleteTodo(ctx, db)
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
