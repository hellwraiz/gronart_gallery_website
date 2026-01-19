package media

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func uploadMedia(c *gin.Context, file *multipart.FileHeader) (string, error) {

	// Generate unique filename
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(file.Filename))

	// Save to disk
	uploadPath := os.Getenv("DATA_DIR") + "images/" + filename
	err := c.SaveUploadedFile(file, uploadPath)
	if err != nil {
		return "", err
	}
	return filename, nil
}
