package media

import (
	"gronart_gallery_website/internal/auth"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type DBHandler struct {
	db *sqlx.DB
}

// Allow the frontend to upload media
// Pls don't exploit these issues 🙏
// TODO: Fix path traversal vulnerability (someone using "../../pwd"), file type validation, file size limit
func create(c *gin.Context) {
	// Get the uploaded file
	file, err := c.FormFile("image")
	if isError(err, "Upload error", http.StatusBadRequest, c) {
		return
	}
	if err = validateImg(file); isError(err, "Upload error", http.StatusBadRequest, c) {
		return
	}

	filename, err := UploadPaintingImg(c, file)
	if isError(err, "Failed to save file", http.StatusInternalServerError, c) {
		return
	}

	// Return the URL
	c.JSON(http.StatusCreated, gin.H{"img_url": filename})
}

func update(c *gin.Context) {

	old_url := c.Param("img_url")
	file, err := c.FormFile("image")
	if isError(err, "Upload error", http.StatusBadRequest, c) {
		return
	}
	if err = validateImg(file); isError(err, "Upload error", http.StatusBadRequest, c) {
		return
	}

	filename, err := UploadPaintingImg(c, file)
	if isError(err, "Failed to save new file", http.StatusInternalServerError, c) {
		return
	}

	if err := DeleteImg(old_url); isError(err, "Failed to delete old image", http.StatusInternalServerError, c) {
		DeleteImg(filename)
		return
	}

	c.JSON(http.StatusOK, gin.H{"img_url": filename})

}

func deleteOne(c *gin.Context) {
	filename := c.Param("img_url")
	if err := DeleteImg(filename); isError(err, "Failed to delete the image", http.StatusInternalServerError, c) {
		return
	}
}

func (h *DBHandler) uploadC(c *gin.Context) {
	db := h.db

	file, err := c.FormFile("image")
	if isError(err, "Upload error", http.StatusBadRequest, c) {
		return
	}
	if err = validateImg(file); isError(err, "Upload error", http.StatusBadRequest, c) {
		return
	}
	filename, err := UploadSiteImg(c, db, "cover", file)
	if isError(err, "DB error", http.StatusInternalServerError, c) {
		return
	}
	c.JSON(http.StatusCreated, gin.H{"cover_url": filename})

}

func (h *DBHandler) deleteC(c *gin.Context) {
	db := h.db

	if err := DeleteSiteImg(c, db, "cover.jpg"); isError(err, "Warning: Failed to delete the cover", http.StatusInternalServerError, c) {
		return
	}
	c.Status(http.StatusAccepted)
}

func (h *DBHandler) updateC(c *gin.Context) {
	db := h.db

	file, err := c.FormFile("image")
	if isError(err, "Upload error", http.StatusBadRequest, c) {
		return
	}
	if err = validateImg(file); isError(err, "Upload error", http.StatusBadRequest, c) {
		return
	}
	if err := DeleteSiteImg(c, db, "cover.jpg"); isError(err, "Warning: Failed to delete the cover", http.StatusInternalServerError, c) {
		return
	}
	filename, err := UploadSiteImg(c, db, "cover", file)
	if isError(err, "DB error", http.StatusInternalServerError, c) {
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"cover_url": filename})
}

func InitRoutes(db *sqlx.DB, api *gin.RouterGroup) {
	media := api.Group("/upload")

	h := DBHandler{db: db}

	// paniting image crud
	media.POST("/", auth.AuthMiddleware(), create)
	media.DELETE("/:img_url", auth.AuthMiddleware(), deleteOne)
	media.PUT("/:img_url", auth.AuthMiddleware(), update)

	// cover image crud
	media.POST("/cover", auth.AuthMiddleware(), h.uploadC)
	media.DELETE("/cover", h.deleteC)
	media.PUT("/cover", auth.AuthMiddleware(), h.updateC)

}
