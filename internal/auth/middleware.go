package auth

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// Middleware to check token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.GetHeader("email")
		pass := c.GetHeader("pass")
		log.Printf("Here's what I got: %s %s\n", email, pass)
		if email == "" || pass == "" {
			c.JSON(401, gin.H{"error": "No token"})
			c.Abort()
			return
		}

		// Check if password is valid
		if os.Getenv("TEMP_EMAIL") == email && os.Getenv("TEMP_PASSWD") == pass {
			c.Set("isLogged", true)
			c.Next()
		} else {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
	}
}
