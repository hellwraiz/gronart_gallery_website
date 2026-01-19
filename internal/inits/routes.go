package inits

import (
	"gronart_gallery_website/internal/media"
	"gronart_gallery_website/internal/paintings"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Initializes all of the routes in this application!
func InitRoutes(db *sqlx.DB) (*gin.Engine, error) {

	// Initiating GIN with a proper env
	router := gin.Default()
	if os.Getenv("GIN_MODE") == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
		router.SetTrustedProxies(nil)
	}

	// Setting up the static routes to be used to server frontend build files
	router.Static("/_app", "./frontend/build/_app")
	router.NoRoute(func(c *gin.Context) {
		c.File("./frontend/build/index.html")
	})
	router.StaticFile("/robots.txt", "./frontend/build/robots.txt")
	// These ones are just to expose the data folder to the frontend though
	router.Static("/images", os.Getenv("DATA_DIR")+"images")

	//// Initiating the api routes
	api := router.Group("/api")

	paintings.InitRoutes(db, api)
	media.InitRoutes(api)

	return router, nil
}
