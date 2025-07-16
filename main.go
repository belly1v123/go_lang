package main

import (
	"go_lang/todo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	g.GET("/", handleInitialRoute)
	g.GET("/todos", todo.HandleGetTodos)

	g.Run(":8080")

}

func handleInitialRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully connected to golan api",
	})

}
