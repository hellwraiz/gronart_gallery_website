package db

import (
	"fmt"

    "github.com/jmoiron/sqlx"		// Makes sql queries take up less space
	_ "github.com/mattn/go-sqlite3" // so that the database/sql package can use sqlite
)

func CreatePainting(db *sqlx.DB, p *Painting) error {
	query := `INSERT INTO paintings (uuid, name, img_url, size, technique) VALUES (?, ?, ?, ?, ?)`

    if _, err := db.Exec(query, generateUUID(), p.Name, p.ImgURL, p.Size, p.Technique); err != nil {
		return fmt.Errorf("Failed to create painting: %s", err)
    }
	return nil
}

func GetPaintingByUUID(db *sqlx.DB, uuid string) (*Painting, error) {

	var p Painting
	p.UUID = uuid

	query := `SELECT name, img_url, size, technique FROM paintings WHERE uuid = ?`


	if err := db.Get(&p, query, uuid); err != nil {
		return nil, fmt.Errorf("Failed to get painting with uuid %s: %s", uuid, err)
	}
	return &p, nil
}

func GetPaintingWithFilter(db *sqlx.DB, filters *Painting, order string) (*Painting, error) {
	return nil, nil

}

func UpdatePainting(db *sqlx.DB, p *Painting) (error) {
	

    query := `
	UPDATE paintings
	SET name = :name, img_url = :img_url, size = :size, technique = :technique, last_edited_at = datetime('now', 'utc')
	WHERE uuid = :uuid
    `
    
    result, err := db.NamedExec(query, p)
    if err != nil {
        return fmt.Errorf("Failed to update painting: %w", err)
    }
	
    // Check if painting exists
    rowsAffected, err := result.RowsAffected()
    if err != nil {
		return fmt.Errorf("Unexpected error occured: %s", err)
    }
    if rowsAffected == 0 {
		return fmt.Errorf("Failed to update painting: painting not found")
    }
    
    return nil
}



