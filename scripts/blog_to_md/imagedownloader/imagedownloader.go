package imagedownloader

import (
	"blog-to-md/models"
	"fmt"
)

// data definitions
// list of images strings[]

var imageUrls []string

// markdown content string
//

func DownloadImages(blogMd models.Blog) (bool, error) {
	fmt.Println("Image processing started...")
	_, err := extractImageUrls(blogMd)

	if err != nil {
		return false, err
	}
	fmt.Println(imageUrls)
	return true, nil
}

func extractImageUrls(blogMd models.Blog) ([]string, error) {
	fmt.Println("Extracting images from .md: ")
	mainImageUrl, err := extractImageFromMetadata()
	//extractImagesFromContent()
	imageUrls = append(imageUrls, mainImageUrl)
	if err != nil {
		return nil, err
	}
	return imageUrls, nil
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
