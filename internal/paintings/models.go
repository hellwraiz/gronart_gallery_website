package paintings

import "time"

type Filter struct {
	Authors    []string
	Sizes      []string
	PriceRange [2]int
	Techniques []string
	Sold       int
	Printable  int
	Copiable   int
	OrderBy    string
	Limit      int
	Offset     int
}

type Painting struct {
	ID           int       `db:"id" json:"-"`
	UUID         string    `db:"uuid" json:"uuid"`
	Name         string    `db:"name" json:"name"`
	Author       string    `db:"author" json:"author"`
	Size         string    `db:"size" json:"size"`
	Price        int       `db:"price" json:"price"`
	ImgURL       string    `db:"img_url" json:"img_url"`
	Technique    string    `db:"technique" json:"technique"`
	Description  string    `db:"description" json:"description"`
	Position     int       `db:"position" json:"position"`
	Sold         int       `db:"sold" json:"sold"`
	Printable    int       `db:"printable" json:"printable"`
	Copiable     int       `db:"copiable" json:"copiable"`
	UploadedAt   time.Time `db:"uploaded_at" json:"uploaded_at"`
	LastEditedAt time.Time `db:"last_edited_at" json:"last_edited_at"`
}

type PatchPainting struct {
	UUID        string  `db:"uuid" json:"uuid"`
	Name        *string `db:"name" json:"name"`
	Author      *string `db:"author" json:"author"`
	Size        *string `db:"size" json:"size"`
	Price       *int    `db:"price" json:"price"`
	ImgURL      *string `db:"img_url" json:"img_url"`
	Technique   *string `db:"technique" json:"technique"`
	Description *string `db:"description" json:"description"`
	Position    *int    `db:"position" json:"position"`
	Sold        *int    `db:"sold" json:"sold"`
	Printable   *int    `db:"printable" json:"printable"`
	Copiable    *int    `db:"copiable" json:"copiable"`
}
