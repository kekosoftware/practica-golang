package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var bookDetailsPattern = regexp.MustCompile(`Title: (.+), Author: (.+), Pages: (\d+)`)

func ParseBooksFromFile(filename string) ([]Book, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var books []Book
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		matches := bookDetailsPattern.FindStringSubmatch(scanner.Text())
		if matches != nil && len(matches) == 4 {
			title := matches[1]
			author := matches[2]
			pages, err := strconv.Atoi(matches[3])
			if err != nil {
				// Log error and continue parsing the rest of the file
				fmt.Printf("Invalid page number for book '%s': %s\n", title, err)
				continue
			}
			books = append(books, Book{Title: title, Author: author, Pages: pages})
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return books, nil
}

func main() {
	filename := "03_book_listings.txt"
	books, err := ParseBooksFromFile(filename)
	if err != nil {
		fmt.Println("Error parsing books from file:", err)
		return
	}
	for _, book := range books {
		fmt.Printf("Parsed Book: %+v\n", book)
	}
}