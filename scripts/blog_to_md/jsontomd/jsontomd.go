package jsontomd

import (
	"blog-to-md/imagedownloader"
	"blog-to-md/models"
	"encoding/json"
	"fmt"
	"log"
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

func (c *JSONToMdConverter) Start() {
	fmt.Println(".json to .md started...")
	createTmpDir(c.jsonFolderName)
	filePaths := findFilePaths(c.jsonFolderName)
	fmt.Println("JSON files to be processed:")
	fmt.Println(filePaths)
	for i := 0; i < len(filePaths); i++ {

		jsonFilePath := fmt.Sprintf("%s/%s", c.jsonFolderName, filePaths[i])
		jsonContent, err := fileToJSON(jsonFilePath)
		imagedownloader.DownloadImages(jsonContent)
		if err != nil {
			log.Fatal("Error converting to JSON:", err)
		}
		newFilePath, err := writeMdFile(c.jsonFolderName, jsonContent)

		if err != nil {
			log.Fatal("Error writing .md: ", err)
		}
		fmt.Println("New file written: ", *newFilePath)
	}

}

func createTmpDir(folderName string) {
	mdFolderName := fmt.Sprintf("%s/mdFiles", folderName)
	if err := os.Mkdir(mdFolderName, 0755); err != nil {
		log.Fatal("Error creating md folder: ", err)
	}
}

func findFilePaths(folderName string) []string {
	var paths []string
	files, err := os.ReadDir(folderName)

	if err != nil {
		log.Fatal("Error reading .json files from folder: ", err)
	}

	for _, entry := range files {
		if !entry.Type().IsDir() {
			paths = append(paths, entry.Name())
		}
	}

	return paths
}

func fileToJSON(path string) (models.Blog, error) {
	var blog models.Blog
	jsonBytes, err := os.ReadFile(path)

	if err != nil {

		log.Fatal("Error reading .json contents:", err)
	}

	if err = json.Unmarshal(jsonBytes, &blog); err != nil {
		log.Fatal("Error parsing json contents: ", err)
	}

	parsedTime, err := time.Parse(time.RFC3339, blog.UpdatedAt)

	if err != nil {
		log.Fatal("Error parsing time string: ", err)
	}

	blog.UpdatedAt = parsedTime.Format("2006-01-02 15:04:05 -0500")

	return blog, nil

}

func writeMdFile(folderPath string, jsonContent models.Blog) (*string, error) {
	fileBytes, err := os.ReadFile("jekyll_example.md")
	if err != nil {
		log.Fatal("Error reading .md example: ", err)
	}
	date := strings.Split(jsonContent.UpdatedAt, " ")[0]
	filename := fmt.Sprintf("%s-%s", date, strings.ReplaceAll(strings.ToLower(jsonContent.Title), " ", "-"))
	newFilePath := fmt.Sprintf("%s/mdFiles/%s.md", folderPath, filename)
	file, err := os.Create(newFilePath)
	if err != nil {
		log.Fatal("Error creating .md:", err)
	}
	defer file.Close()

	t := template.Must(template.New("template").Parse(string(fileBytes)))
	err = t.Execute(file, jsonContent)

	if err != nil {
		return nil, err
	}
	return &newFilePath, nil
}
