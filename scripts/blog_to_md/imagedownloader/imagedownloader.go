package imagedownloader

import (
	"blog-to-md/models"
	"fmt"
)

func DownloadImages(jsonContent models.Blog) {
	var images []string
	fmt.Println("Image processing started...")

	images = append(images, jsonContent.MainImage)
	fmt.Println(images)
}
