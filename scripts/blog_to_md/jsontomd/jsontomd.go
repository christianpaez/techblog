package jsontomd

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type JSONToMdConverter struct {
	jsonFolderName string
}

type Blog struct {
	Title   string `json:"title"`
	Content string `json:content`
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

		jsonContent, err := fileToJSON(filePaths[i])
		if err != nil {
			log.Fatal("Error converting to JSON:", err)
		}
		fmt.Println(jsonContent)

		_, err = writeMdFile(jsonContent)

		if err != nil {
			log.Fatal("Error writing .md: ", err)
		}
		fmt.Println("File written: ", filePaths[i])
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
	// todo

	fmt.Println(path)

	return Blog{
		Title:   "Some title",
		Content: "Some content",
	}, nil

}

func writeMdFile(jsonContent Blog) (bool, error) {
	// todo
	const testTemplate = `Im am template, some value: {{.Title}} and {{.Content}}`
	t := template.Must(template.New("testTemplate").Parse(testTemplate))
	err := t.Execute(os.Stdout, jsonContent)

	if err != nil {
		return false, err
	}
	return true, nil
}
