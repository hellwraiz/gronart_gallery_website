package paintings

//// Setting up all of the crud operations. TODO: ADD A LOOOT OF DATA VALIDATION, and identity verification for some of these
import (
	"gronart_gallery_website/internal/auth"
	"gronart_gallery_website/internal/media"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type DBHandler struct {
	db *sqlx.DB
}

func (h *DBHandler) create(c *gin.Context) {
	var painting Painting
	db := h.db
	if err := c.BindJSON(&painting); isError(err, "JSON error", http.StatusBadRequest, c) {
		return
	}
	if err := CreatePainting(db, &painting); isError(err, "DB error", http.StatusInternalServerError, c) {
		return
	}
	c.JSON(http.StatusCreated, painting)
}

func (h *DBHandler) getFiltered(c *gin.Context) {
	db := h.db
	// Getting the parameters
	authors, sizes, priceRangeStr, techniques, soldStr, printableStr, copiableStr, orderBy, limitStr, offsetStr :=
		c.QueryArray("authors"), c.QueryArray("sizes"), c.QueryArray("price_range"),
		c.QueryArray("techniques"), c.Query("sold"), c.Query("printable"),
		c.Query("copiable"), c.Query("order_by"), c.Query("limit"), c.Query("offset")

	// Doing data processing/validation
	// TODO: improve data validation
	var priceRange [2]int
	log.Printf("Here's the price range %v", priceRangeStr)
	if len(priceRangeStr) == 2 {
		priceRange = [2]int{StoI(priceRangeStr[0], -1), StoI(priceRangeStr[1], -1)}
	} else {
		log.Printf("Couldn't parse this price range: %v", priceRangeStr)
		priceRange = [2]int{-1, -1}
	}

	sold := StoI(soldStr, -1)
	printable := StoI(printableStr, -1)
	copiable := StoI(copiableStr, -1)
	limit := StoI(limitStr, -1)
	offset := StoI(offsetStr, -1)

	// Populate the filter thing
	filters := &Filter{Authors: authors, Sizes: sizes, PriceRange: priceRange,
		Techniques: techniques, Sold: sold, Printable: printable, Copiable: copiable,
		OrderBy: orderBy, Limit: limit, Offset: offset}

	// Get the actual paintings!
	paintings, err := GetPaintingWithFilter(db, filters)
	if isError(err, "DB error", http.StatusInternalServerError, c) {
		return
	}

	c.JSON(http.StatusOK, paintings)
}

func (h *DBHandler) deleteOne(c *gin.Context) {
	db := h.db
	// get the painting, and img_url
	uuid := c.Param("uuid")
	painting, err := GetPaintingByUUID(db, uuid)
	if isError(err, "DB error", http.StatusInternalServerError, c) {
		return
	}

	imgUrl := painting.ImgURL

	if err = DeletePainting(db, uuid); isError(err, "DB error", http.StatusInternalServerError, c) {
		return
	}

	if err := media.DeleteImg(imgUrl); err != nil {
		log.Printf("Warning: failed to delete uploaded file: %s", err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "Painting deleted"})

}

func (h *DBHandler) patch(c *gin.Context) {
	db := h.db

	var painting PatchPainting
	uuid := c.Param("uuid")
	if err := c.ShouldBindJSON(&painting); isError(err, "JSON error", http.StatusBadRequest, c) {
		return
	}
	painting.UUID = uuid
	newPainting, err := UpdatePainting(db, &painting)
	if isError(err, "DB error", http.StatusInternalServerError, c) {
		return
	}
	c.JSON(http.StatusOK, newPainting)
}

func InitRoutes(db *sqlx.DB, api *gin.RouterGroup) {
	paintings := api.Group("/paintings")

	h := DBHandler{db: db}

	paintings.GET("", h.getFiltered)
	paintings.POST("", auth.AuthMiddleware(), h.create)
	paintings.DELETE("/:uuid", auth.AuthMiddleware(), h.deleteOne)
	paintings.PATCH("/:uuid", auth.AuthMiddleware(), h.patch)

}
