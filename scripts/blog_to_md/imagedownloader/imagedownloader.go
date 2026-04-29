package imagedownloader

import (
	"blog-to-md/models"
	"fmt"
)

// data definitions
// list of images strings[]

// markdown content string
//

func DownloadImages(jsonContent models.Blog) error {
	var images []string
	fmt.Println("Image processing started...")

	images = append(images, jsonContent.MainImage)
	fmt.Println(images)
	return nil
}

func extractImageUrls() {
	fmt.Println("Extracting images from .md: ")
	extractImageFromMetadata()
	extractImagesFromContent()
}

func extractImageFromMetadata() (string, error) {
	fmt.Println("Extracting from metadata")
	return "ifojeioj", nil
}

func extractImagesFromContent() ([]string, error) {
	imageUrls := []string{"1", "2", "3"}
	return imageUrls, nil
}

func downloadImagesToFs() error {
	return nil
}
