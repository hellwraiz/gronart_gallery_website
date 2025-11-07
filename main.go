package main

import (
	"gronart_gallery_website/internal/db"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"      // the very nice library for routing
	"github.com/joho/godotenv"      // gives me access to the .env file values in the app
	_ "github.com/mattn/go-sqlite3" // so that the database/sql package can use sqlite
)

func main() {
	
	// loading .env
	godotenv.Load()

	// starting the database
	database, err := db.InitDB()
	if err != nil {
		log.Fatal("Couldn't initiate the database:", err)
	}
    defer database.Close()

	// Initiating GIN
	router := gin.Default()
	if os.Getenv("GIN_MODE") == "release" {
		router.SetTrustedProxies(nil)
	}
	
	// Setting up default routes
	router.Static("/assets", "./frontend/dist/assets")
	router.NoRoute(func(c *gin.Context) {
		c.File("./frontend/dist/index.html")
	})

	// Setting up all the routes
	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Starting the server
    port := os.Getenv("PORT")
	log.Println("Starting server on port", port)
	router.Run()
}
