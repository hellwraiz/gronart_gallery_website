package users

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	tokenStr := c.GetHeader("Authorization")
	// strip "Bearer " prefix
	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

	claims, err := verifyToken(tokenStr)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("user_id", claims["user_id"])
	c.Set("role", claims["role"])
	c.Next()
}

func AdminMiddleware(c *gin.Context) {
	if c.MustGet("role") == "admin" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	c.Next()
}
