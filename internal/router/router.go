package router

import (
	"codebase-golang/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/users", handler.GetUsers)

	// r.GET("/excel", handler.FetchExcel)

	return r
}
