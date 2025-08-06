package handler

import (
	"net/http"

	"codebase-golang/internal/service"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users, err := service.FetchUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
