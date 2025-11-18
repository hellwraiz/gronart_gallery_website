package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"       // Makes sql queries take up less space
	_ "github.com/mattn/go-sqlite3" // so that the database/sql package can use sqlite
)

type Filter struct {
	Authors			[]string
    Sizes        	[]string
	PriceRange		[]int
    Techniques   	[]string
	OrderBy			string
	Limit			int
	Offset			int
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

	query := `SELECT * FROM paintings WHERE 1=1`

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

	if filters.PriceRange != nil {
		query += " AND "
		if filters.PriceRange[0] != -1 {
			if filters.PriceRange[1] != -1 {
				query += "price BETWEEN ? AND ?"
				args = append(args, filters.PriceRange[0], filters.PriceRange[1])
			} else {
				query += "price >= ?"
				args = append(args, filters.PriceRange[0])
			}
		} else if filters.PriceRange[1] != -1 {
			query += "price <= ?"
			args = append(args, filters.PriceRange[1])
		}
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

	if filters.OrderBy != "" {
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

func UpdatePainting(db *sqlx.DB, p *Painting) (error) {
	
    query := `
	UPDATE paintings
	SET name = :name, author = :author, size = :size, price = :price, img_url = :img_url, technique = :technique, last_edited_at = datetime('now', 'utc')
	WHERE uuid = :uuid
    `
    
    result, err := db.NamedExec(query, p)
    if err != nil {
        return fmt.Errorf("Failed to update painting: %w", err)
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


