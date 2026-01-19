package inits

import (
	"fmt"
	"gronart_gallery_website/internal/paintings"
	"log"
	"os"

	"github.com/jmoiron/sqlx" // Makes sql queries take up less space
)

func InitDB() (*sqlx.DB, error) {
	// initializing the connection
	dataDir := os.Getenv("DATA_DIR")
	dbPath := dataDir + "gallery.db"
	db, err := sqlx.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	log.Printf("Database connected successfully on path: %s", dbPath)

	// Testing if the database was actually created.
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("Failed to connect to the database: %s", err)
	}

	err = paintings.InitDB(db)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("Failed to initiate paintings database: %s", err)
	}

	// alles goed!
	return db, nil
}
