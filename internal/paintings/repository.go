package paintings

import (
	"fmt"

	"github.com/jmoiron/sqlx"       // Makes sql queries take up less space
	_ "github.com/mattn/go-sqlite3" // so that the database/sql package can use sqlite
)

func CreatePainting(db *sqlx.DB, p *Painting) error {

	query := `INSERT INTO paintings (uuid, name, author, size, price, img_url, technique, description, sold, printable, copiable) VALUES (:uuid, :name, :author, :size, :price, :img_url, :technique, :description, :sold, :printable, :copiable)`

	p.UUID = generateUUID()

	// TODO: refactor all code to look like this. Very nice
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

// TODO: Change so we don't need to use args array. Use named thing
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

	if filters.Sold != false {
		query += " AND sold = ?"
		args = append(args, filters.Sold)
	}

	if filters.Printable != false {
		query += " AND printable = ?"
		args = append(args, filters.Printable)
	}

	if filters.Copiable != false {
		query += " AND copiable = ?"
		args = append(args, filters.Copiable)
	}

	if filters.OrderBy != "" {
		allowedOrders := map[string]bool{"price": true, "name": true, "uploaded_at": true, "position": true}
		if allowedOrders[filters.OrderBy] {
			query += " ORDER BY ?"
			args = append(args, filters.OrderBy)
		} else {
			return nil, fmt.Errorf("Invalid order operation.")
		}
	} else {
		query += " ORDER BY position"
	}

	if filters.Limit != -1 {
		query += " LIMIT ?"
		args = append(args, filters.Limit)
		if filters.Offset != -1 {
			query += " OFFSET ?"
			args = append(args, filters.Offset)
		}
	}

	query += ";"

	if err := db.Select(&paintings, query, args...); err != nil {
		return nil, fmt.Errorf("Failed to get paintings with filters %v: %s", filters, err)
	}
	return &paintings, nil
}

func UpdatePainting(db *sqlx.DB, p *PatchPainting) (*Painting, error) {

	query := `
	UPDATE paintings
	SET 
    `

	if p.Name != nil {
		query += "name = :name, "
	}

	if p.Author != nil {
		query += "author = :author, "
	}

	if p.Size != nil {
		query += "size = :size, "
	}

	if p.Price != nil {
		query += "price = :price, "
	}

	if p.ImgURL != nil {
		query += "img_url = :img_url, "
	}

	if p.Technique != nil {
		query += "technique = :technique, "
	}

	if p.Description != nil {
		query += "description = :description, "
	}

	if p.Position != nil {
		query += "position = :position, "
	}

	if p.Sold != nil {
		query += "sold = :sold, "
	}

	if p.Printable != nil {
		query += "printable = :printable, "
	}

	if p.Copiable != nil {
		query += "copiable = :copiable, "
	}

	query += `last_edited_at = datetime('now', 'utc')
	WHERE uuid = :uuid;`

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

func DeletePainting(db *sqlx.DB, position int, uuid string) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("Failed to start transaction: %s", err)
	}

	query := `
	UPDATE paintings
	SET position = position - 1
	WHERE position > ?;
	`

	result, err := tx.Exec(query, position)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Failed to delete painting: %s", err)
	}
	// not checking for rows affected in case it was the last row

	query = `
	DELETE FROM paintings
	WHERE uuid = ?;
	`

	result, err = tx.Exec(query, uuid)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Failed to delete painting: %s", err)
	}
	if err := rowsAffected(result); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func ReorderPaintings(db *sqlx.DB, r *Reordering) error {

	// Putting the painting at index 0
	tx, err := db.Beginx()
	if err != nil {
		return fmt.Errorf("Failed to start transaction: %s", err)
	}
	query := `
	UPDATE paintings
	SET position = 0
	WHERE position = :source;`

	result, err := tx.NamedExec(query, r)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Failed to reorder paintings: %s", err)
	}
	if err := rowsAffected(result); err != nil {
		tx.Rollback()
		return err
	}

	if r.Source > r.Destination {
		query = `
		UPDATE paintings
		SET position = position + 1
		WHERE position >= :destination and position < :source;`
	} else {
		query = `
		UPDATE paintings
		SET position = position - 1
		WHERE position > :source and position <= :destination;`
	}
	result, err = tx.NamedExec(query, r)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Failed to reorder paintings: %s", err)
	}
	if err := rowsAffected(result); err != nil {
		tx.Rollback()
		return err
	}

	// Moving the source painting to target position
	query = `
	UPDATE paintings
	SET position = :destination
	WHERE position = 0;
	`

	result, err = tx.NamedExec(query, r)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Failed to reorder paintings: %s", err)
	}
	if err := rowsAffected(result); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
