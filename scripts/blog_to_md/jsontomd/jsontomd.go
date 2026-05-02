package jsontomd

import (
	"blog-to-md/imagedownloader"
	"blog-to-md/models"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"
)

type JSONToMdConverter struct {
	jsonFolderName string
}

func NewJSONToMdConverter(jsonFolderName string) *JSONToMdConverter {
	c := JSONToMdConverter{jsonFolderName: jsonFolderName}
	return &c
}

func (c *JSONToMdConverter) Start() (bool, error) {
	fmt.Println(".json to .md started...")
	if err := createTmpDir(c.jsonFolderName); err != nil {
		return false, err
	}
	filePaths, err := findFilePaths(c.jsonFolderName)
	if err != nil {
		return false, err
	}
	fmt.Println("JSON files to be processed:")
	fmt.Println(filePaths)
	for i := 0; i < len(filePaths); i++ {
		jsonFilePath := fmt.Sprintf("%s/%s", c.jsonFolderName, filePaths[i])
		jsonContent, err := fileToJSON(jsonFilePath)
		if err != nil {
			return false, err
		}
		imageDownloader := imagedownloader.NewDownloader(jsonContent)
		if _, err := imageDownloader.DownloadImages(); err != nil {
			return false, err
		}
		newFilePath, err := writeMdFile(c.jsonFolderName, jsonContent)
		if err != nil {
			return false, err
		}
		fmt.Println("New file written: ", *newFilePath)
	}
	return true, nil
}

func createTmpDir(folderName string) error {
	mdFolderName := fmt.Sprintf("%s/mdFiles", folderName)
	return os.Mkdir(mdFolderName, 0755)
}

func findFilePaths(folderName string) ([]string, error) {
	var paths []string
	files, err := os.ReadDir(folderName)
	if err != nil {
		return nil, err
	}
	for _, entry := range files {
		if !entry.Type().IsDir() {
			paths = append(paths, entry.Name())
		}
	}
	return paths, nil
}

func fileToJSON(path string) (models.Blog, error) {
	var blog models.Blog
	jsonBytes, err := os.ReadFile(path)
	if err != nil {
		return blog, err
	}
	if err = json.Unmarshal(jsonBytes, &blog); err != nil {
		return blog, err
	}
	parsedTime, err := time.Parse(time.RFC3339, blog.UpdatedAt)
	if err != nil {
		return blog, err
	}
	blog.UpdatedAt = parsedTime.Format("2006-01-02 15:04:05 -0500")
	return blog, nil
}

func writeMdFile(folderPath string, jsonContent models.Blog) (*string, error) {
	fileBytes, err := os.ReadFile("jekyll_example.md")
	if err != nil {
		return nil, err
	}
	date := strings.Split(jsonContent.UpdatedAt, " ")[0]
	filename := fmt.Sprintf("%s-%s", date, strings.ReplaceAll(strings.ToLower(jsonContent.Title), " ", "-"))
	newFilePath := fmt.Sprintf("%s/mdFiles/%s.md", folderPath, filename)
	file, err := os.Create(newFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	t := template.Must(template.New("template").Parse(string(fileBytes)))
	err = t.Execute(file, jsonContent)
	if err != nil {
		return nil, err
	}
	return &newFilePath, nil
}
