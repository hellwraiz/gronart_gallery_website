package media

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadImg(c *gin.Context, file *multipart.FileHeader) (string, error) {

	// HACK: Temporarily let this add a cover.jpg image that allows you to set some image as the cover
	// Generate unique filename
	var filename string
	if file.Filename == "cover.jpg" {
		filename = "cover.jpg"
	} else {
		filename = fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(file.Filename))
	}

	// Save to disk
	uploadPath := os.Getenv("DATA_DIR") + "images/" + filename
	err := c.SaveUploadedFile(file, uploadPath)
	if err != nil {
		return "", err
	}
	return filename, nil
}

func DeleteImg(filename string) error {

	err := os.Remove(os.Getenv("DATA_DIR") + "images/" + filename)
	if err != nil {
		return fmt.Errorf("Warning: failed to delete uploaded file: %s", err)
	}

	return nil

}
