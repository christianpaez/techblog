package main

import (
	"blog-to-md/inputprompt"
	"blog-to-md/jsontomd"
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-sigChan
		fmt.Println("\n Interrupted, closing...")
		os.Exit(1)
	}()

	inputprompt.Wait("Program started, press enter to continue...")
	credentials, errTest := inputprompt.Command("Press enter to check credentials file...", func() ([]byte, error) {
		return os.ReadFile(".credentials.txt")
	})

	handleErr(errTest)

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
	fileNumber := 1
	folderName := fmt.Sprintf("tmp/%s", time.Now().Format("2006-01-02_150405.000000"))

	if _, err := os.Stat("tmp"); os.IsNotExist(err) {
		if err := os.Mkdir("tmp", 0755); err != nil {
			log.Fatal("Error creating tmp directory:", err)
		}
	}
	if err := os.Mkdir(folderName, 0755); err != nil {
		log.Fatal("Error creating tmp nested directory:", err)
	}
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
				fileName := fmt.Sprintf("0%s-%s.json", strconv.Itoa(fileNumber), time.Now().Format("2006-01-02_150405.000000"))
				path := filepath.Join(folderName, fileName)
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
				fmt.Printf("Wrote %d bytes\n", fileBytes)

				w.Flush()

				fileNumber += 1
			}
		})

	}

	fmt.Println("Files Created!")

	inputprompt.Command("Press enter to format files to JSON...", jsontomd.NewJSONToMdConverter(folderName).Start)
}
