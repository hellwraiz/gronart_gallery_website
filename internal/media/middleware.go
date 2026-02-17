package media

import (
	"github.com/gin-gonic/gin"
)

// Middleware to check that the image is correct
func ImageChecking() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
