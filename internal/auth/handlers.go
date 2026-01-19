package auth

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	var login Login
	c.BindJSON(&login)

	// Check if password is valid
	if os.Getenv("TEMP_EMAIL") == login.Email && os.Getenv("TEMP_PASSWD") == login.Password {
		c.Status(http.StatusOK)
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
	}
}

func Init(api *gin.RouterGroup) {
	auth := api.Group("/login")

	auth.POST("/login", login)

}
