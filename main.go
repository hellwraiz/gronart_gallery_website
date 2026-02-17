package main

import (
	"fmt"
	"gronart_gallery_website/internal/inits"
	_ "gronart_gallery_website/internal/inits"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	_ "log"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"    // gives me access to the .env file values in the app
	_ "github.com/mattn/go-sqlite3" // so that the database/sql package can use sqlite
)

func main() {

	// loading .env. Only necessary if in dev!
	// Otherwise fly does this functions job. Since no .env file in prod
	if os.Getenv("GIN_MODE") != gin.ReleaseMode {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file: ", err)
		}
	}

	mainn()

	// starting the db
	db, err := inits.InitDB()
	if err != nil {
		log.Fatal("Couldn't initiate the database: ", err)
	}
	defer db.Close()

	// Initiating the routes
	router, err := inits.InitRoutes(db)
	if err != nil {
		log.Fatal("Couldn't initiate the routes: ", err)
	}

	// Starting the server. It automatically gets the environment port variable.
	router.Run()
}

func mainn() {
	imagesDir := os.Getenv("DATA_DIR") + "images/"

	files, err := os.ReadDir(imagesDir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filename := file.Name()

		// Skip if already a thumbnail
		if strings.Contains(filename, "_thumb") {
			continue
		}
		// Skip if it's a cover image
		if strings.Contains(filename, "cover") {
			continue
		}

		// Check if thumbnail already exists
		ext := filepath.Ext(filename)
		thumbFilename := strings.TrimSuffix(filename, ext) + "_thumb" + ext
		thumbPath := imagesDir + thumbFilename

		if _, err := os.Stat(thumbPath); err == nil {
			fmt.Println("Thumbnail exists, skipping:", filename)
			continue
		}

		// Generate thumbnail
		err := createThumbnailForFile(imagesDir+filename, thumbPath)
		if err != nil {
			fmt.Printf("Failed to create thumbnail for %s: %v\n", filename, err)
			continue
		}

		fmt.Println("Created thumbnail for:", filename)
	}

	fmt.Println("Done!")
}

func createThumbnailForFile(originalPath, thumbPath string) error {
	// Open original
	f, err := os.Open(originalPath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Decode
	img, format, err := image.Decode(f)
	if err != nil {
		return err
	}

	// Resize
	thumbnail := resize.Resize(400, 0, img, resize.Lanczos3)

	// Create thumbnail file
	out, err := os.Create(thumbPath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Encode
	if format == "png" {
		return png.Encode(out, thumbnail)
	}
	return jpeg.Encode(out, thumbnail, &jpeg.Options{Quality: 85})
}
