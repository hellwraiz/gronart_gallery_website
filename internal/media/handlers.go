package media

import (
	"gronart_gallery_website/internal/auth"
	"log"
	"net/http"
	"os"

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

	filename, err := UploadImg(c, file)
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

	filename, err := UploadImg(c, file)
	if isError(err, "Failed to save new file", http.StatusInternalServerError, c) {
		return
	}

	if err := DeleteImg(old_url); isError(err, "Failed to delete old image", http.StatusInternalServerError, c) {
		DeleteImg(filename)
		return
	}

	c.JSON(http.StatusOK, gin.H{"img_url": filename})

}

func deleteC(c *gin.Context) {

	err := os.Remove(os.Getenv("DATA_DIR") + "images/cover.jpg")
	if err != nil {
		log.Printf("Warning: failed to delete uploaded file: %s\n", err)
	}
	c.Status(http.StatusAccepted)
}

func deleteOne(c *gin.Context) {
	filename := c.Param("img_url")
	if err := DeleteImg(filename); isError(err, "Failed to delete the image", http.StatusInternalServerError, c) {
		return
	}
}

func InitRoutes(api *gin.RouterGroup) {
	media := api.Group("/upload")

	media.POST("/", auth.AuthMiddleware(), create)
	media.GET("/delete", deleteC)
	media.DELETE("/:img_url", auth.AuthMiddleware(), deleteOne)
	media.PUT("/:img_url", auth.AuthMiddleware(), update)

}
