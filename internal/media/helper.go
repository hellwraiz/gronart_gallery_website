package media

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	_ "golang.org/x/image/webp"

	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
)

const MaxImageSize = 8 * 1024 * 1024 // 8MB

// ---------------------------- handler helpers ----------------------------
func isError(err error, errorMsg string, errorCode int, c *gin.Context) bool {
	if err != nil {
		err := fmt.Sprintf("%s: %s", errorMsg, err)
		log.Println(err)
		c.JSON(errorCode, gin.H{"error": err})
		return true
	}
	return false
}

func validateImg(file *multipart.FileHeader) error {

	if file.Size > MaxImageSize {
		return fmt.Errorf("image too large, max %dMB", MaxImageSize/1028/1028)
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	buffer := make([]byte, 512)
	_, err = src.Read(buffer)
	if err != nil {
		return err
	}

	mimeType := http.DetectContentType(buffer)
	switch mimeType {
	case "image/jpeg", "image/png", "image/webp":
		return nil
	default:
		return fmt.Errorf("unsupported format, use JPG, WebP or PNG")
	}
}

// -------------------------- repository helpers ---------------------------
func CreateThumbnailForFile(originalPath, thumbPath string) error {
	f, err := os.Open(originalPath)
	if err != nil {
		return err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		f.Close()
		return err
	}

	thumbnail := resize.Resize(400, 0, img, resize.Lanczos3)

	thumbPath = strings.TrimSuffix(thumbPath, filepath.Ext(thumbPath)) + ".jpg"

	out, err := os.Create(thumbPath)
	if err != nil {
		return err
	}
	defer out.Close()

	return jpeg.Encode(out, thumbnail, &jpeg.Options{Quality: 85})
}
