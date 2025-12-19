package database

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
	Price			int			`db:"price" json:"price"`
    ImgURL      	string    	`db:"img_url" json:"img_url"`
    Technique   	string    	`db:"technique" json:"technique"`
    UploadedAt  	time.Time   `db:"uploaded_at" json:"uploaded_at"`
    LastEditedAt 	time.Time   `db:"last_edited_at" json:"last_edited_at"`
}

type User struct {
    ID           string    `json:"id" gorm:"primaryKey"`
    Email        string    `json:"email" gorm:"unique;not null"`
    PasswordHash string    `json:"-" gorm:"not null"` // Never send to frontend
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}

// Optional: Session table (if using stateful tokens)
type Session struct {
    Token     string    `gorm:"primaryKey"`
    UserID    string    `gorm:"not null"`
    ExpiresAt time.Time `gorm:"not null"`
    CreatedAt time.Time
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
		price INTEGER NOT NULL,
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
