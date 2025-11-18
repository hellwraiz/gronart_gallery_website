package main

import (
	"gronart_gallery_website/internal/db"
	"gronart_gallery_website/internal/routes"
	"log"
	"os"

	"github.com/joho/godotenv"      // gives me access to the .env file values in the app
	_ "github.com/mattn/go-sqlite3" // so that the database/sql package can use sqlite
)

func main() {
	
	// loading .env
	godotenv.Load()

	// starting the database
	database, err := db.InitDB()
	if err != nil {
		log.Fatal("Couldn't initiate the database: ", err)
	}
    defer database.Close()

	// Initiating the routes
	router, err := routes.InitRoutes()
	if err != nil {
		log.Fatal("Couldn't initiate the routes: ", err)
	}

	// Starting the server
    port := os.Getenv("PORT")
	log.Println("Starting server on port", port)
	router.Run()
	
}
