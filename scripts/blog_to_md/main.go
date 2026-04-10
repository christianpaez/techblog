package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
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
		s := []string{"https://dev.to", value[1], "/edit"}
		url := fmt.Sprintf(strings.Join(s, ""))
		fmt.Printf("Visiting blog page: %s\n", url)

		client := &http.Client{}

		currentUserCookie := &http.Cookie{
			Name:   "current_user",
			Value:  "",
			Quoted: false,
		}

		rememberMeCookie := &http.Cookie{
			Name:   "remember_user_token",
			Value:  "",
			Quoted: false,
		}

		request, err := http.NewRequest("GET", url, nil)

		if err != nil {

			log.Fatal("Error creating request: ", err)
		}

		request.AddCookie(currentUserCookie)
		request.AddCookie(rememberMeCookie)
		request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

		request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
		request.Header.Set("Accept-Language", "en-US,en;q=0.5")
		request.Header.Set("Connection", "keep-alive")
		request.Header.Set("Upgrade-Insecure-Requests", "1")
		response, err := client.Do(request)

		if err != nil {
			panic(err)
		}

		defer response.Body.Close()

		fmt.Printf("%v\n", response.Status)

		document, err := goquery.NewDocumentFromReader(response.Body)

		if err != nil {
			log.Fatal(err)

		}
		document.Find("#main-content").Each(func(i int, selection *goquery.Selection) {

			attribute, exists := selection.Attr("data-article")

			if exists {
				fmt.Printf("Attribute: %s \n", attribute)
			}
		})

	}

}
