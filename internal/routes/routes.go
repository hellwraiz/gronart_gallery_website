package routes

import (
	"fmt"
	"gronart_gallery_website/internal/database"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

// Initializes all of the routes in this application!
func InitRoutes(db *sqlx.DB) (*gin.Engine, error) {

	godotenv.Load()

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
	// These ones are just to expose the data folder to the frontend
	router.Static("/images", os.Getenv("DATA_DIR")+"images")

	// Setting up the api routes
	{
		api := router.Group("/api")

		router.GET("/ping", AuthMiddleware(), func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		//// Setting up all of the crud operations. TODO: ADD A LOOOT OF DATA VALIDATION, and identity verification for some of these
		// Get filtered paintings
		api.GET("/paintings", func(c *gin.Context) {
			// Getting the parameters
			authors, sizes, priceRangeStr, techniques, orderBy, limitStr, offsetStr :=
				c.QueryArray("authors"), c.QueryArray("sizes"), c.QueryArray("price_range"),
				c.QueryArray("techniques"), c.Query("order_by"), c.Query("limit"), c.Query("offset")

			// Doing data processing/validation
			// TODO: improve data validation
			var priceRange [2]int
			if len(priceRangeStr) == 2 {
				priceRange = [2]int{StoI(priceRangeStr[0], -1), StoI(priceRangeStr[1], -1)}
			} else {
				log.Printf("Couldn't parse this price range: %v", priceRangeStr)
				priceRange = [2]int{-1, -1}
			}

			limit := StoI(limitStr, 20)
			offset := StoI(offsetStr, 0)

			// Populate the filter thing
			filters := &database.Filter{Authors: authors, Sizes: sizes, PriceRange: priceRange,
				Techniques: techniques, OrderBy: orderBy, Limit: limit, Offset: offset}

			// Get the actual paintings!
			paintings, err := database.GetPaintingWithFilter(db, filters)
			if isError(err, "DB error", http.StatusInternalServerError, c) {
				return
			}

			c.JSON(http.StatusOK, paintings)
		})

		// Create a painting
		api.POST("/paintings", AuthMiddleware(), func(c *gin.Context) {
			var painting database.Painting
			if err := c.BindJSON(&painting); isError(err, "JSON error", http.StatusBadRequest, c) {
				return
			}
			if err := database.CreatePainting(db, &painting); isError(err, "DB error", http.StatusInternalServerError, c) {
				return
			}
			c.JSON(http.StatusCreated, painting)
		})

		// Allow the frontend to upload media
		// Pls don't exploit these issues 🙏
		// TODO: Fix path traversal vulnerability (someone using "../../pwd"), file type validation, file size limit
		api.POST("/upload", AuthMiddleware(), func(c *gin.Context) {
			// Get the uploaded file
			file, err := c.FormFile("image")
			if isError(err, "Upload error", http.StatusBadRequest, c) {
				return
			}

			// Generate unique filename
			filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(file.Filename))

			// Save to disk
			uploadPath := os.Getenv("DATA_DIR") + "images/" + filename
			if err = c.SaveUploadedFile(file, uploadPath); isError(err, "Failed to save file", http.StatusInternalServerError, c) {
				return
			}

			// Return the URL
			c.JSON(http.StatusCreated, gin.H{"img_url": filename})
		})

		api.DELETE("/paintings", AuthMiddleware(), func(c *gin.Context) {
			// get the painting, and img_url
			uuid := c.Query("uuid")
			painting, err := database.GetPaintingByUUID(db, uuid)
			if isError(err, "DB error", http.StatusInternalServerError, c) {
				return
			}

			imgUrl := painting.ImgURL

			if err = database.DeletePainting(db, uuid); isError(err, "DB error", http.StatusInternalServerError, c) {
				return
			}

			err = os.Remove(os.Getenv("DATA_DIR") + "images/" + imgUrl)
			if err != nil {
				log.Printf("Warning: failed to delete uploaded file: %s", err)
			}
			c.JSON(http.StatusOK, gin.H{"message": "Painting deleted"})
		})

		api.POST("/login", func(c *gin.Context) {
			var login Login
			c.BindJSON(&login)

			// Check if password is valid
			if os.Getenv("TEMP_EMAIL") == login.Email && os.Getenv("TEMP_PASSWD") == login.Password {
				c.Status(http.StatusOK)
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			}
		})
	}

	return router, nil
}
