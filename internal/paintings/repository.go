package paintings

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"       // Makes sql queries take up less space
	_ "github.com/mattn/go-sqlite3" // so that the database/sql package can use sqlite
)

func InitDB(db *sqlx.DB) error {

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

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	query = `CREATE INDEX IF NOT EXISTS idx_paintings_uuid ON paintings(uuid);`

	_, err = db.Exec(query)

	return err

}

func CreatePainting(db *sqlx.DB, p *Painting) error {

	query := `INSERT INTO paintings (uuid, name, author, size, price, img_url, technique) VALUES (:uuid, :name, :author, :size, :price, :img_url, :technique)`

	p.UUID = generateUUID()

	if result, err := db.NamedExec(query, p); err != nil {
		return fmt.Errorf("Failed to create painting: %s", err)
	} else if numAffected, errResult := result.RowsAffected(); numAffected == 0 && errResult == nil {
		return fmt.Errorf("Failed to create painting: Database unaffected")
	} else if errResult != nil {
		return fmt.Errorf("Couldn't find out if painting was created: %s", errResult)
	}
	return nil
}

func GetPaintingByUUID(db *sqlx.DB, uuid string) (*Painting, error) {

	var p Painting
	p.UUID = uuid

	query := `SELECT * FROM paintings WHERE uuid = ?`

	if err := db.Get(&p, query, uuid); err != nil {
		return nil, fmt.Errorf("Failed to get painting with uuid %s: %s", uuid, err)
	}
	return &p, nil
}

func GetPaintingWithFilter(db *sqlx.DB, filters *Filter) (*[]Painting, error) {

	var paintings []Painting
	var args []any

	query := "SELECT * FROM paintings WHERE 1=1"

	if filters.Authors != nil {
		query += " AND "
		for index, author := range filters.Authors {
			if index == 0 {
				query += "author = ?"
				args = append(args, author)
			} else {
				query += " OR author = ?"
				args = append(args, author)
			}
		}
	}

	if filters.Sizes != nil {
		query += " AND "
		for index, size := range filters.Sizes {
			if index == 0 {
				query += "size = ?"
				args = append(args, size)
			} else {
				query += " OR size = ?"
				args = append(args, size)
			}
		}
	}

	if filters.PriceRange[0] != -1 {
		query += " AND "
		if filters.PriceRange[1] != -1 {
			query += "price BETWEEN ? AND ?"
			args = append(args, filters.PriceRange[0], filters.PriceRange[1])
		} else {
			query += "price >= ?"
			args = append(args, filters.PriceRange[0])
		}
	} else if filters.PriceRange[1] != -1 {
		query += " AND price <= ?"
		args = append(args, filters.PriceRange[1])
	}

	if filters.Techniques != nil {
		query += " AND "
		for index, technique := range filters.Sizes {
			if index == 0 {
				query += "technique = ?"
				args = append(args, technique)
			} else {
				query += " OR technique = ?"
				args = append(args, technique)
			}
		}
	}

	if filters.OrderBy != "" { // sql injection attack possibility!!!!! Match against allowed thingies
		// will use this later on
		/* allowedOrders := map[string]bool{"price": true, "name": true, "uploaded_at": true}
		   if allowedOrders[*filters.OrderBy] {
		       fb.query.WriteString(" ORDER BY " + *filters.OrderBy)
		   } */
		query += " ORDER BY ?"
		args = append(args, filters.OrderBy)
	}

	if filters.Limit != 0 {
		query += " LIMIT ?"
		args = append(args, filters.Limit)
		if filters.Offset != 0 {
			query += " OFFSET ?"
			args = append(args, filters.Offset)
		}
	}

	log.Println("here's the final query:", query)

	query += ";"

	if err := db.Select(&paintings, query, args...); err != nil {
		return nil, fmt.Errorf("Failed to get paintings with filters %v: %s", filters, err)
	}
	return &paintings, nil
}

func UpdatePainting(db *sqlx.DB, p *PostPainting) (*Painting, error) {

	var args []any

	/* query := `
		UPDATE paintings
		SET name = :name, author = :author, size = :size, price = :price, img_url = :img_url, technique = :technique, last_edited_at = datetime('now', 'utc')
		WHERE uuid = :uuid
	    ` */

	query := `
	UPDATE paintings
	SET 
    `

	if p.Name != nil {
		query += "name = :name, "
		args = append(args, *p.Name)
	}

	if p.Author != nil {
		query += "author = :author, "
		args = append(args, *p.Author)
	}

	if p.Size != nil {
		query += "size = :size, "
		args = append(args, *p.Size)
	}

	if p.Price != nil {
		query += "price = :price, "
		args = append(args, *p.Price)
	}

	if p.ImgURL != nil {
		query += "img_url = :img_url, "
		args = append(args, *p.ImgURL)
	}

	if p.Technique != nil {
		query += "technique = :technique, "
		args = append(args, *p.Technique)
	}

	query += `last_edited_at = datetime('now', 'utc')
	WHERE uuid = :uuid`

	result, err := db.NamedExec(query, p)
	if err != nil {
		return nil, fmt.Errorf("Failed to update painting: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("Unexpected error occured: %s", err)
	}
	if rowsAffected == 0 {
		return nil, fmt.Errorf("Failed to update painting: painting not found")
	}

	newPainting, err := GetPaintingByUUID(db, p.UUID)
	if err != nil {
		return nil, fmt.Errorf("Failed to get updated painting: %s", err)
	}
	return newPainting, nil
}

func DeletePainting(db *sqlx.DB, uuid string) error {
	query := `
	DELETE FROM paintings
	WHERE uuid = ? 
	`

	result, err := db.Exec(query, uuid)
	if err != nil {
		return fmt.Errorf("Failed to delete painting: %s", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Unexpected error occured: %s", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("Failed to update painting: painting not found")
	}

	return nil
}
