package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println("Start program")
	fmt.Println("Reading file...")

	contents, err := os.ReadFile("index_example.html")

	if err != nil {
		log.Printf("Error reading file: %v", err)
	}
	fmt.Printf("File contents: %s", contents)
}
