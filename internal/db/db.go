package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"      // gives me access to the .env file values in the app
	_ "github.com/mattn/go-sqlite3" // so that the database/sql package can use sqlite
)

type Painting struct {
    ID          int       `json:"-"` // keep the id internal. No need to send it to ever send it to the frontend
    UUID        string    `json:"uuid"`
    Name        string    `json:"name"`
    Size        string    `json:"size"`
    ImgURL      string    `json:"img_url"`
    Technique   string    `json:"technique"`
    UploadedAt  string    `json:"uploaded_at"`
    LastEditedAt string   `json:"last_edited_at"`
}

func InitDB() (*sql.DB, error) {
	// loading .env
	godotenv.Load()

	// initializing the connection
	dataDir := os.Getenv("DATA_DIR")
	dbPath := filepath.Join(dataDir, "gallery.db")
    db, err := sql.Open("sqlite3", dbPath)
    if err != nil {
		return nil, err
    }
	log.Printf("Database connected successfully on path: %s", dbPath)
	log.Printf("Check this shit out %s", filepath.Join("data", "gallery.db"))
	log.Printf("And now this shit %s", filepath.Join("/data", "gallery.db"))
	log.Printf("Now this shit %s", filepath.Join("/data/", "gallery.db"))
	log.Printf("This shit %s", filepath.Join("data/", "gallery.db"))

	// Just checking if it works yknow
    if err := db.Ping(); err != nil {
        db.Close()
        return nil, err
    }

	// initializing the actual database
	query := `
	CREATE TABLE IF NOT EXISTS paintings (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid TEXT NOT NULL UNIQUE,
		name TEXT NOT NULL,
		img_url TEXT NOT NULL,
		size TEXT NOT NULL,
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

func CreatePainting(db *sql.DB, p *Painting) error {
	query := `INSERT INTO paintings (uuid, name, img_url, size, technique) VALUES (?, ?, ?, ?, ?)`

    if _, err := db.Exec(query, p.UUID, p.Name, p.ImgURL, p.Size, p.Technique); err != nil {
        return err
    } else {
		return nil
	}
}

func GetPaintingByUUID(db *sql.DB, uuid string) (*Painting, error) {

	var p Painting
	p.UUID = uuid

	query := `SELECT name, img_url, size, technique FROM paintings WHERE uuid = ?`

	if err := db.QueryRow(query, uuid).Scan(&p.Name, &p.ImgURL, &p.Size, &p.Technique); err != nil {
		return nil, err
	} else {
		return &p, err
	}
}
