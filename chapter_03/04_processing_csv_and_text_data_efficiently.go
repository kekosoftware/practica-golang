package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func ImportBooksFromCSV(filename string) ([]Book, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var books []Book
	for _, record := range records {
		pages, err := strconv.Atoi(record[2])
		if err != nil {
			// Handle error
			continue
		}
		books = append(books, Book{
			Title:  record[0],
			Author: record[1],
			Pages:  pages,
		})
	}
	return books, nil
}
func ExportBooksToCSV(filename string, books []Book) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, book := range books {
		record := []string{book.Title, book.Author, strconv.Itoa(book.Pages)}
		if err := writer.Write(record); err != nil {
			// Handle error
			return err
		}
	}
	return nil
}
func main() {
	filename := "books.csv"
	// Assume books is populated with Book structs
	if err := ExportBooksToCSV(filename, books); err != nil {
		fmt.Printf("Failed to export books to CSV: %s\n", err)
	}
	importedBooks, err := ImportBooksFromCSV(filename)
	if err != nil {
		fmt.Printf("Failed to import books from CSV: %s\n", err)
	} else {
		fmt.Println("Imported Books:", importedBooks)
	}
}