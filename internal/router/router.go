package router

import (
	"github.com/gin-gonic/gin"

	handler "codebase-golang/internal/handler"
	"codebase-golang/internal/service"
)

func NewRouter(userSvc *service.UserService) *gin.Engine {
	r := gin.Default()

	userHandler := handler.NewUserHandler(userSvc)

	api := r.Group("/api")
	{
		api.GET("/users", userHandler.GetAllUsers)
		api.POST("/users/refresh", userHandler.RefreshUsers) // manual refresh
	}

	return r
}
