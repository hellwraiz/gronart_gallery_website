package routes

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

// Middleware to check token
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
		var login Login
        /* authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(401, gin.H{"error": "No token"})
            c.Abort()
            return
        } */
		c.BindJSON(&login)
		log.Printf("Here's what I got: %v\n", login)
        
        
        // Check if password is valid
		log.Printf("Here's my current env: %s %s\n", os.Getenv("TEMP_EMAIL"), os.Getenv("TEMP_PASSWD"))
		if os.Getenv("TEMP_EMAIL") == login.Email && os.Getenv("TEMP_PASSWD") == login.Password {
			c.Set("isLogged", true)
			c.Next()
		} else {
            c.JSON(401, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }
    }
}
