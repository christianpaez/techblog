package jsontomd

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
)

type JSONToMdConverter struct {
	jsonFolderName string
}

type Blog struct {
	Title     string `json:"title"`
	Content   string `json:"body_markdown"`
	UpdatedAt string `json:"updated_at"`
	MainImage string `json:"main_image"`
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

func fileToJSON(path string) (Blog, error) {
	var blog Blog
	jsonBytes, err := os.ReadFile(path)

	if err != nil {

		log.Fatal("Error reading .json contents:", err)
	}

	if err = json.Unmarshal(jsonBytes, &blog); err != nil {
		log.Fatal("Error parsing json contents: ", err)
	}

	return blog, nil

}

func writeMdFile(folderPath string, jsonContent Blog) (*string, error) {
	const testTemplate = `Im am template, some value: {{.Title}} and {{.Content}}`
	newFilePath := fmt.Sprintf("%s/mdFiles/%s.md", folderPath, jsonContent.UpdatedAt)
	file, err := os.Create(newFilePath)
	if err != nil {
		log.Fatal("Error creating .md:", err)
	}
	defer file.Close()

	t := template.Must(template.New("testTemplate").Parse(testTemplate))
	err = t.Execute(file, jsonContent)

	if err != nil {
		return nil, err
	}
	return &newFilePath, nil
}
