package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "auth will be implemented in lab 4",
		"token":   "stub",
	})
}

func (h *Handlers) Logout(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func (h *Handlers) Register(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.Auth.Register(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}
