package main

import (
	"github.com/joho/godotenv"
	// "github.com/mattn/go-sqlite3"
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"os"
	// "log"
)

func main() {
	
	godotenv.Load()
    port := os.Getenv("PORT")
    if port == "" {
		fmt.Println("failed to scan port from .env. Defaulting to 8080")
        port = "8080"
    }
	router := gin.Default()
	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	fmt.Println("Starting server on port", port)
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
