package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	fmt.Println("Start program")
	fmt.Println("Reading file...")

	credentials, err := os.ReadFile((".credentials.txt"))

	scanner := bufio.NewScanner(strings.NewReader(string(credentials)))

	var currentUserToken, rememberUserToken string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "current_user:") {
			currentUserToken = strings.TrimPrefix(line, "current_user:")
		}

		if strings.HasPrefix(line, "remember_user_token:") {
			rememberUserToken = strings.TrimPrefix(line, "remember_user_token:")
		}
	}
	if currentUserToken == "" || rememberUserToken == "" {
		log.Fatal("Credentials missing. Check .credentials.txt format.")
	}
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
			Value:  currentUserToken,
			Quoted: false,
		}

		rememberMeCookie := &http.Cookie{
			Name:   "remember_user_token",
			Value:  rememberUserToken,
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
				fmt.Printf("Writing to file...: %s \n", attribute)
				path := filepath.Join("tmp/", "check")
				f, err := os.Create(path)
				if err != nil {
					log.Fatal("Error creating file: ", err)

				}
				defer f.Close()
				w := bufio.NewWriter(f)
				fileBytes, err := w.WriteString(attribute)
				if err != nil {

					log.Fatal("Error writing found attribute to file: ", err)
				}

				fmt.Printf("Wrote %d bytes", fileBytes)

			}
		})

	}

}
