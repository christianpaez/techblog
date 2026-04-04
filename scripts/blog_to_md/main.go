package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {

	fmt.Println("Start program")
	fmt.Println("Reading file...")

	contents, err := os.ReadFile("index_example.html")

	if err != nil {
		log.Printf("Error reading file: %v", err)
	}
	fmt.Printf("File contents: %s", contents)

	validHrefRegex := regexp.MustCompile("href=\"([^\"]+)\"")
	matches := validHrefRegex.FindAllStringSubmatch(string(contents), -1)
	for _, value := range matches {
		fmt.Printf("Match group found: %s\n", value[1])
	}

}
