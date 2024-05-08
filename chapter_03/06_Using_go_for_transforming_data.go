package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func GenerateLibrarySummary(books []Book) {
	fmt.Printf("Total Books: %d\n", len(books))
	var totalPages int
	booksByAuthor := make(map[string][]Book)
	for _, book := range books {
		totalPages += book.Pages
		booksByAuthor[book.Author] = append(booksByAuthor[book.Author], book)
	}
	avgPages := float64(totalPages) / float64(len(books))
	fmt.Printf("Average Pages per Book: %.2f\n", avgPages)
	for author, books := range booksByAuthor {
		fmt.Printf("%s has %d books\n", author, len(books))
	}
}
func ExportLibraryDataForAnalysis(filename string, books []Book) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	// Write header
	if err := writer.Write([]string{"Title", "Author", "Pages"}); err != nil {
		return err
	}
	// Write book data
	for _, book := range books {
		if err := writer.Write([]string{book.Title, book.Author, strconv.Itoa(book.Pages)}); err != nil {
			return err
		}
	}
	return nil
}
func main() {
	// Assuming books is a slice of Book populated with the user's library data
	GenerateLibrarySummary(books)
	if err := ExportLibraryDataForAnalysis("library_analysis.csv", books); err != nil {
		fmt.Printf("Failed to export library data for analysis: %s\n", err)
	}
}
