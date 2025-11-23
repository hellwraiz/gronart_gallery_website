package main

import (
	"gronart_gallery_website/internal/database"
	"gronart_gallery_website/internal/routes"
	"log"

	"github.com/joho/godotenv"      // gives me access to the .env file values in the app
	_ "github.com/mattn/go-sqlite3" // so that the database/sql package can use sqlite
)

func main() {
	
	// loading .env
	godotenv.Load()

	// starting the db
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Couldn't initiate the database: ", err)
	}
    defer db.Close()

	// Initiating the routes
	router, err := routes.InitRoutes(db)
	if err != nil {
		log.Fatal("Couldn't initiate the routes: ", err)
	}

	// Starting the server. It automatically gets the environment port variable.
	log.Print("This is before starting the servenr")
	router.Run()
	log.Print("This is after starting the server")
	
}
