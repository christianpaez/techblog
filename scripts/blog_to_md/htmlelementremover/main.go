package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Need filename")
		return
	}
	filename := os.Args[1]
	html, err := os.ReadFile(filename)

	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(html)))

	doc.Find(".dashboard-story__actions").Remove()
	doc.Find(".pagination").Remove()
	cleaned, _ := doc.Html()
	os.WriteFile("index_example.html", []byte(cleaned), 0644)
}
