package main

import (
	"gronart_gallery_website/internal/inits"
	"log"

	"github.com/joho/godotenv"      // gives me access to the .env file values in the app
	_ "github.com/mattn/go-sqlite3" // so that the database/sql package can use sqlite
)

func main() {

	// loading .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// starting the db
	db, err := inits.InitDB()
	if err != nil {
		log.Fatal("Couldn't initiate the database: ", err)
	}
	defer db.Close()

	// Initiating the routes
	router, err := inits.InitRoutes(db)
	if err != nil {
		log.Fatal("Couldn't initiate the routes: ", err)
	}

	// Starting the server. It automatically gets the environment port variable.
	router.Run()
}
