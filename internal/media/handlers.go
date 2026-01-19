package media

import (
	"gronart_gallery_website/internal/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Allow the frontend to upload media
// Pls don't exploit these issues 🙏
// TODO: Fix path traversal vulnerability (someone using "../../pwd"), file type validation, file size limit
func create(c *gin.Context) {
	// Get the uploaded file
	file, err := c.FormFile("image")
	if isError(err, "Upload error", http.StatusBadRequest, c) {
		return
	}

	filename, err := uploadMedia(c, file)
	if isError(err, "Failed to save file", http.StatusInternalServerError, c) {
		return
	}

	// Return the URL
	c.JSON(http.StatusCreated, gin.H{"img_url": filename})

}

func InitRoutes(api *gin.RouterGroup) {
	media := api.Group("/upload")

	media.POST("", auth.AuthMiddleware(), create)

}
