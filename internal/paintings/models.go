package paintings

import "time"

type Filter struct {
	Authors    []string
	Sizes      []string
	PriceRange [2]int
	Techniques []string
	OrderBy    string
	Limit      int
	Offset     int
}

// The schema
// If you change this, also change the query in InitDB, as well as everything in crud.go
type Painting struct {
	ID           int       `db:"id" json:"-"`
	UUID         string    `db:"uuid" json:"uuid"`
	Name         string    `db:"name" json:"name"`
	Author       string    `db:"author" json:"author"`
	Size         string    `db:"size" json:"size"`
	Price        int       `db:"price" json:"price"`
	ImgURL       string    `db:"img_url" json:"img_url"`
	Technique    string    `db:"technique" json:"technique"`
	UploadedAt   time.Time `db:"uploaded_at" json:"uploaded_at"`
	LastEditedAt time.Time `db:"last_edited_at" json:"last_edited_at"`
}

type PostPainting struct {
	UUID      string  `json:"uuid" binding:"required"`
	Name      *string `json:"name"`
	Author    *string `json:"author"`
	Size      *string `json:"size"`
	Price     *int    `json:"price"`
	ImgURL    *string `json:"img_url"`
	Technique *string `json:"technique"`
}
