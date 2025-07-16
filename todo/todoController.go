package todo

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGetTodos(ctx *gin.Context) {
	Alltodos := GetDummyTodos()
	ctx.JSON(http.StatusOK, Alltodos)
}

func HandleAddTodos(ctx *gin.Context) {
	var todo Todo
	err := ctx.ShouldBindJSON(&todo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	todo.ID = len(GetDummyTodos()) + 1 // Simple ID assignment logic
	log.Println("Adding Todo:", todo)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Todo added successfully", "todo": todo})
}
