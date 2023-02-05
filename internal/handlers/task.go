package handlers

import (
	"net/http"
	"taskboard/internal/service"

	"github.com/gin-gonic/gin"
)

func GetTasks(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, service.FetchTasks())
}
