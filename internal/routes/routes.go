package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitRoutes() (*gin.Engine, error) {

	godotenv.Load()

	log.Print("This is before setting the environment variable")
	// Initiating GIN with a proper env
	router := gin.Default()
	if os.Getenv("GIN_MODE") == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
		router.SetTrustedProxies(nil)
	}
	log.Print("This is after setting the environment variable")

	// Setting up the static routes to be used to server frontend build files
	router.Static("/assets", "./frontend/dist/assets")
	router.NoRoute(func(c *gin.Context) {
		c.File("./frontend/dist/index.html")
	})

	// Setting up the api routes
	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return router, nil
}
