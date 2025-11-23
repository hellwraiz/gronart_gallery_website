package routes

import (
	"gronart_gallery_website/internal/database"
	"log"
	"net/http"
	"os"
	"path/filepath"

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


	//// Setting up all of the crud operations. TODO: ADD A LOOOT OF DATA VALIDATION
	// Get filtered paintings
	router.GET("/api/paintings", func(c *gin.Context) {
		// Getting the parameters
		authors := c.QueryArray("authors")
		sizes := c.QueryArray("sizes")
		priceRangeStr := c.QueryArray("price_range")
		techniques := c.QueryArray("techniques")
		orderBy := c.Query("order_by")
		limitStr := c.Query("limit")
		offsetStr := c.Query("offset")

		// Doing data processing/validation
		// TODO: improve data validation
		priceRange := [2]int{StoI(priceRangeStr[0], -1), StoI(priceRangeStr[1], -1)}

		limit := StoI(limitStr, 20)
		offset := StoI(offsetStr, 0)

		// Populate the filter thing
		filters := &database.Filter{
			Authors: authors,
			Sizes: sizes,
			PriceRange: priceRange,
			Techniques: techniques,
			OrderBy: orderBy,
			Limit: limit,
			Offset: offset,
		}

		// Get the actual paintings!
		paintings, err := database.GetPaintingWithFilter(db, filters)
		if err != nil {
			log.Printf("DB error: %s", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch paintings",
			})
			return
		}

		c.JSON(http.StatusOK, paintings)
	})

	// Create a painting
	router.POST("/api/paintings", func(c *gin.Context) {
		var painting database.Painting

		if err := c.BindJSON(&painting); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON format",
			})
			return
		}

		if err := database.CreatePainting(db, &painting); err != nil {
			log.Printf("DB error: %s", err)
		}
	})

	// Allow the frontend to upload media
	router.POST("/api/upload", func(c *gin.Context) {
		// Get the uploaded file
		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
			return
		}

		// Generate unique filename
		filename := filepath.Ext(file.Filename)

		// Save to disk
		uploadPath := os.Getenv("DATA_DIR") + filename
		if err := c.SaveUploadedFile(file, uploadPath); err != nil {
			log.Printf("Failed to save file: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}

		// Return the URL
		c.JSON(http.StatusCreated, gin.H{
			"img_url": uploadPath, 
		})
	})



	return router, nil
}
