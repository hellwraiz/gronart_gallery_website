package database

import (
	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3" // so that the database/sql package can use sqlite
)

func generateUUID() string {
    return uuid.New().String()
}

