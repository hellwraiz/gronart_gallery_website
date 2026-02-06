package inits

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	// Running migrations
	if err := runMigrations(db); err != nil {
		db.Close()
		return nil, err
	}

	// alles goed!
	return db, nil
}

func runMigrations(db *sqlx.DB) error {
	driver, err := sqlite3.WithInstance(db.DB, &sqlite3.Config{}) // gives golang-migrate access to db
	if err != nil {
		return fmt.Errorf("Failed to create migration driver: %s", err)
	}

	m, err := migrate.NewWithDatabaseInstance( // creates database instance
		"file://migrations", // tells where the migration files are
		"sqlite3",           // tells what kind of database to use
		driver,              // tells which DB to connect to
	)
	if err != nil {
		return fmt.Errorf("Failed to initiate migration driver: %s", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange { // the part that actually does the migrations
		return fmt.Errorf("Failed to perform the migrations: %s", err)
	}

	return nil
}
