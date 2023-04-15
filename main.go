package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Get up", Completed: false},
	{ID: "2", Item: "Dress up", Completed: false},
	{ID: "3", Item: "Go to college", Completed: false},
}

func getTodods(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func main() {
	router := gin.Default()
	router.GET("/todos")
	router.Run("localhost:9090")
}
