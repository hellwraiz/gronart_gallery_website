package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"       // Makes sql queries take up less space
	"github.com/joho/godotenv"      // gives me access to the .env file values in the app
	_ "github.com/mattn/go-sqlite3" // so that the database/sql package can use sqlite
)

// The schema
// If you change this, also change the query in InitDB, as well as everything in crud.go
type Painting struct {
	ID				int			`db:"id" json:"-"`
    UUID        	string    	`db:"uuid" json:"uuid"`
    Name        	string    	`db:"name" json:"name"`
	Author			string		`db:"author" json:"author"`
    Size        	string    	`db:"size" json:"size"`
	Price			string		`db:"price" json:"price"`
    ImgURL      	string    	`db:"img_url" json:"img_url"`
    Technique   	string    	`db:"technique" json:"technique"`
    UploadedAt  	time.Time   `db:"uploaded_at" json:"uploaded_at"`
    LastEditedAt 	time.Time   `db:"last_edited_at" json:"last_edited_at"`
}

func InitDB() (*sqlx.DB, error) {
	// loading .env
	godotenv.Load()

	// initializing the connection
	dataDir := os.Getenv("DATA_DIR")
	dbPath := dataDir + "gallery.db"
    db, err := sqlx.Open("sqlite3", dbPath)
    if err != nil {
		return nil, err
    }
	log.Printf("Database connected successfully on path: %s", dbPath)

	// Just checking if it works yknow
    if err := db.Ping(); err != nil {
        db.Close()
		return nil, fmt.Errorf("Failed to connect to the database: %s", err)
    }

	// initializing the actual database
	query := `
	CREATE TABLE IF NOT EXISTS paintings (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid TEXT NOT NULL UNIQUE,
		name TEXT NOT NULL,
		author TEXT NOT NULL,
		size TEXT NOT NULL,
		price TEXT NOT NULL,
		img_url TEXT NOT NULL,
		technique TEXT NOT NULL,
		uploaded_at DATETIME DEFAULT (datetime('now', 'utc')),
		last_edited_at DATETIME DEFAULT (datetime('now', 'utc'))
	);
	`
    
	_, err = db.Exec(query)
	if err != nil {
		db.Close()
		return nil, err
	}

	query = `CREATE INDEX idx_paintings_uuid ON paintings(uuid);`

	// alles goed!
    return db, nil
}
