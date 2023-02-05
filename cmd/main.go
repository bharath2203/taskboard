package main

import (
	"taskboard/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/tasks", handlers.GetTasks)

	r.Run("localhost:8080")
}
