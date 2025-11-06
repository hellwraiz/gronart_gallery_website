package main

import (
	"github.com/joho/godotenv"
	// "github.com/mattn/go-sqlite3"
	"github.com/gin-gonic/gin"
	// "fmt"
	"net/http"
	"os"
	"log"
)

func main() {
	
	godotenv.Load()
    port := os.Getenv("PORT")
	
	log.Println("Application starting...")
	log.Printf("Environment: %s", os.Getenv("ENV"))
    if port == "" {
		log.Println("failed to scan port from .env. Defaulting to 8080")
        port = "8080"
    }
	router := gin.Default()
	
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
	log.Println("Starting server on port", port)
	router.Run()

	/* // Serve static files from Svelte build
	http.Handle("/", http.FileServer(http.Dir("./frontend/dist")))

	// API routes
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"message": "Hello from Go!"}`)
	})

	fmt.Println("Starting server on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil)) */
}
