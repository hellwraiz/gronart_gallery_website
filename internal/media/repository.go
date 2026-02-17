package media

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/nfnt/resize"
)

func UploadPaintingImg(c *gin.Context, file *multipart.FileHeader) (string, error) {

	// Generate unique filenames
	ext := filepath.Ext(file.Filename)
	uploadTime := time.Now().UnixNano()
	filename := fmt.Sprintf("%d%s", uploadTime, ext)
	thumbName := fmt.Sprintf("%d_thumb%s", uploadTime, ext)

	// Save original image
	uploadPath := os.Getenv("DATA_DIR") + "images/" + filename
	err := c.SaveUploadedFile(file, uploadPath)
	if err != nil {
		return "", err
	}

	// Open image
	f, err := os.Open(uploadPath)
	if err != nil {
		DeleteImg(filename)
		return "", fmt.Errorf("Thumbnail err: %s", err)
	}
	defer f.Close()

	// Resize the image
	img, format, err := image.Decode(f)
	if err != nil {
		DeleteImg(filename)
		return "", fmt.Errorf("Thumbnail err: %s", err)
	}
	thumbnail := resize.Resize(400, 0, img, resize.Lanczos3)

	// Converting and saving thumbnail to disk
	uploadPath = os.Getenv("DATA_DIR") + "images/" + thumbName
	out, err := os.Create(uploadPath)
	if err != nil {
		DeleteImg(filename)
		out.Close()
		return "", fmt.Errorf("Thumbnail err: %s", err)
	}
	defer out.Close()

	if format == "png" {
		err = png.Encode(out, thumbnail)
		if err != nil {
			DeleteImg(filename)
			return "", fmt.Errorf("Thumbnail err: %s", err)
		}
	} else {
		err = jpeg.Encode(out, thumbnail, &jpeg.Options{Quality: 85})
		if err != nil {
			DeleteImg(filename)
			return "", fmt.Errorf("Thumbnail err: %s", err)
		}
	}

	return filename, nil

}

func DeleteImg(filename string) error {

	if filename == "" {
		return fmt.Errorf("No image url uploaded")
	}
	err := os.Remove(os.Getenv("DATA_DIR") + "images/" + filename)
	if err != nil {
		return fmt.Errorf("Warning: failed to delete uploaded file: %s", err)
	}

	return nil

}

func UploadSiteImg(c *gin.Context, db *sqlx.DB, key string, file *multipart.FileHeader) (string, error) {
	ext := filepath.Ext(file.Filename)
	uploadPath := os.Getenv("DATA_DIR") + "images/" + key + ext
	err := c.SaveUploadedFile(file, uploadPath)
	if err != nil {
		return "", err
	}

	query := "INSERT INTO site_config (key, value) VALUES (?, ?)"
	if result, err := db.Exec(query, key, key+ext); err != nil {
		return "", fmt.Errorf("Failed to upload site %s: %s", key, err)
	} else if numAffected, errResult := result.RowsAffected(); numAffected == 0 && errResult == nil {
		return "", fmt.Errorf("Failed to upload site %s: Database unaffected", key)
	} else if errResult != nil {
		return "", fmt.Errorf("Couldn't find out if site %s was created: %s", key, errResult)
	}
	return "cover" + ext, nil
}

func DeleteSiteImg(c *gin.Context, db *sqlx.DB, key string) error {
	query := `
	DELETE FROM site_config
	WHERE key = ? 
	`

	result, err := db.Exec(query, key)
	if err != nil {
		return fmt.Errorf("Failed to delete %s: %s", key, err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Unexpected error occured: %s", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("Failed to update %s: %s not found", key, key)
	}

	return nil

}
