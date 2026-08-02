package imagedownloader

import (
	"blog-to-md/config"
	"blog-to-md/models"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"slices"
	"strings"
)

type Downloader struct {
	imageUrls []string
	blogMd    models.Blog
}

func NewDownloader(blogMg models.Blog) *Downloader {
	downloader := Downloader{blogMd: blogMg}
	return &downloader
}

func (downloader *Downloader) DownloadImages() (bool, error) {
	fmt.Println("Image processing started...")
	err := downloader.extractImageUrls()

	if err != nil {
		return false, err
	}
	return true, nil
}

func (downloader *Downloader) extractImageUrls() error {
	fmt.Println("Extracting images from .md: ")
	mainImageUrl, err := extractImageFromMetadata(downloader.blogMd)
	contentImages, err := extractImagesFromContent(downloader.blogMd)
	downloader.imageUrls = slices.Concat(append(downloader.imageUrls, mainImageUrl), contentImages)
	createImageDir()
	err = downloader.downloadImagesToFs()
	if err != nil {
		return err
	}
	return nil
}

func extractImageFromMetadata(blogMd models.Blog) (string, error) {
	return blogMd.MainImage, nil
}

func extractImagesFromContent(blogMd models.Blog) ([]string, error) {
	reg, err := regexp.Compile(`https?://[^\s]+\.(png|jpg|jpeg|gif|webp)`)

	if err != nil {
		return nil, err
	}
	imageUrls := reg.FindAllString(blogMd.Content, 2)
	return imageUrls, nil
}

func (downloader Downloader) downloadImagesToFs() error {
	for i := 0; i < len(downloader.imageUrls); i++ {
		fmt.Printf("Downloading image: %s\n", downloader.imageUrls[i])
		downloader.downloadImage(i)
	}
	return nil
}

func createImageDir() error {
	return os.Mkdir(config.ImageDir, 0755)
}

func (downloader *Downloader) downloadImage(index int) error {
	imagePathPrefix := blogTitleToImagePrefix(downloader.blogMd.Title)
	response, err := http.Get(downloader.imageUrls[index])

	if err != nil {
		return err
	}

	defer response.Body.Close()

	imageFilePath := fmt.Sprintf("%s/%s-0%d.png", config.ImageDir, imagePathPrefix, index)
	file, err := os.Create(imageFilePath)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, response.Body)

	if err != nil {
		return err
	}
	return nil
}

func blogTitleToImagePrefix(title string) string {
	return strings.ReplaceAll(strings.ToLower(title), " ", "-")

}

func (downloader *Downloader) GetImageUrls() []string {
	return downloader.imageUrls
}
