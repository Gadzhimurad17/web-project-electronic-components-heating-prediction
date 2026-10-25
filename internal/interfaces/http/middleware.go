package http

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type middleware struct {
	secretKey []byte
}

const prefix = "Bearer"

func (m *middleware) extractToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, prefix) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "error"})
		return "", fmt.Errorf("")
	}

	token := strings.TrimSpace(strings.TrimPrefix(authHeader, prefix))
	if token == "" || strings.ContainsAny(token, " \t\r\n") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token format",
		})
		return "", fmt.Errorf("")

	}
	return token, nil
}

func (m *middleware) validateToken(c *gin.Context) error {
	JWTtoken, err := m.extractToken(c)
	if err != nil {
		return err
	}

	token, err := jwt.Parse(JWTtoken, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok != true {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return m.secretKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return fmt.Errorf("token was expired (TTL истек)")
		}
		return fmt.Errorf("failed to parse token: %w", err)
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return nil
}

func (m *middleware) StrictAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := m.validateToken(c); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.Next()
	}
}

func (m *middleware) LenientAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := m.validateToken(c); err != nil {
			return
		}
		c.Next()
	}
}
