package jekyllmdpathhelper

import (
	"blog-to-md/config"
	"fmt"
	"os"
)

// data definitions
// string need config md path where files are located
//
// string|template need header string to insert when content does not have header image
//
// string need a blog with md content
//
// string[] for each blog need list of image paths for lookup and replace

type JekyllMdPathHelper struct {
	contentUrls []string
}

func NewJekyllMdPathHelper(contentUrls []string) *JekyllMdPathHelper {
	return &JekyllMdPathHelper{
		contentUrls: contentUrls,
	}
}

func (jekyllMdPathHelper *JekyllMdPathHelper) Normalize() error {
	// stub this reads from md folder and calls helper
	files, err := os.ReadDir(config.GetMdDir())
	if err != nil {
		return err
	}
	jekyllMdPathHelper.replacePaths(files)
	fmt.Println(jekyllMdPathHelper.contentUrls)
	return nil
}

func insertHeader() {} // stub takes image path and inserts header withmain image url

func replaceImagePath() {} // stub replace path with new path

func (jekyllmDPathHelper *JekyllMdPathHelper) replacePaths(paths []os.DirEntry) error {
	for _, entry := range paths {
		// this is wrong, i need to assume i know file path and the content urls will be used in constructor so i can do a single loop
		if !entry.IsDir() {
			filePath := fmt.Sprintf("%s/%s", config.GetMdDir(), entry.Name())
			fileBytes, err := os.ReadFile(filePath)
			if err != nil {
				return err
			}

			fmt.Println(string(fileBytes))
		}
	}
	return nil
}
