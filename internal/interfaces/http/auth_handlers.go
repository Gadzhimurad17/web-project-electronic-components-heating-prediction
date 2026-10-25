package http

import (
	"heating-prediction/internal/infrastructure/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) Login(c *gin.Context) {
	req := dto.LoginRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func (h *Handlers) Logout(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func (h *Handlers) Register(c *gin.Context) {
	req := dto.LoginRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user, err := h.Auth.Register(c.Request.Context(), req.Login, req.Password)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}
