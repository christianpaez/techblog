package jsontomd

import "fmt"

type blog struct {
	title   string `json:"title"`
	content string `json:content`
}

func Start() {
	fmt.Println(".json to .md started...")
	createTmpDir()
	filePaths := findFilePaths()
	fmt.Println(filePaths)
	filesToJSON()
	//fmt.Println(jsonContents)

}

func createTmpDir() {
	// todo
}

func findFilePaths() []string {
	// todo
	paths := []string{"g", "h", "i"}
	return paths
}

func filesToJSON() {
	// todo
	// var jsonContents blog
	// ill continue here but its late, gotta do a loop
	// and marshall file data to json and them a go template
	// to get .mds...
}
