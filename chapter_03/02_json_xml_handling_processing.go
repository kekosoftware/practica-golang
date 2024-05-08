package main

import (
	"encoding/xml"
	"fmt"
)

type Book struct {
	Title  string `json:"title" xml:"title"`
	Author string `json:"author" xml:"author"`
	Pages  int    `json:"pages" xml:"pages"`
}

// For XML, we often work with a wrapper type to represent a collection of books.
type Library struct {
	Books []Book `xml:"book"`
}

func ExportBooksToXML(books []Book) (string, error) {
	library := Library{Books: books}
	xmlData, err := xml.MarshalIndent(library, "", " ")
	if err != nil {
		return "", err
	}
	return string(xmlData), nil
}
func ImportBooksFromXML(xmlData string) ([]Book, error) {
	var library Library
	err := xml.Unmarshal([]byte(xmlData), &library)
	if err != nil {
		return nil, err
	}
	return library.Books, nil
}

func main() {
	books := []Book{
		{"The Go Programming Language", "Alan A. A. Donovan", 380},
		{"Go in Action", "William Kennedy", 300},
	}
	xmlOutput, err := ExportBooksToXML(books)
	if err != nil {
		fmt.Println("Error exporting books to XML:", err)
		return
	}
	fmt.Println("XML Output:", xmlOutput)
	// Simulate importing books from XML
	importedBooks, err := ImportBooksFromXML(xmlOutput)
	if err != nil {
		fmt.Println("Error importing books from XML:", err)
		return
	}
	fmt.Println("Imported Books:", importedBooks)
}