package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
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
		s := []string{"https://dev.to", value[1]}
		url := fmt.Sprintf(strings.Join(s, ""))
		fmt.Printf("Visiting blog page: %s\n", url)
		response, err := http.Get(url)

		if err != nil {
			panic(err)
		}

		defer response.Body.Close()

		fmt.Printf("%v\n", response.Status)

	}

}
