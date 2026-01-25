package auth

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	var login Login
	c.BindJSON(&login)

	// Check if password is valid
	log.Printf("here's what I use %s %s", os.Getenv("TEMP_EMAIL"), os.Getenv("TEMP_PASSWD"))
	if os.Getenv("TEMP_EMAIL") == login.Email && os.Getenv("TEMP_PASSWD") == login.Password {
		c.Status(http.StatusOK)
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
	}
}

func InitRoutes(api *gin.RouterGroup) {
	auth := api.Group("/login")

	auth.POST("", login)

}
