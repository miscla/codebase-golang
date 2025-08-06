package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"codebase-golang/internal/service"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.svc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"source": "cache_or_db", "data": users})
}

func (h *UserHandler) RefreshUsers(c *gin.Context) {
	if err := h.svc.RefreshCache(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "refreshed"})
}
