package main

import (
	"fmt"
	"io"
	"os"
)

type Book struct {
	Title     string
	Author    string
	Pages     int
	CoverPath string // Path to cover image file
}

func ReadCoverImage(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	imageData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return imageData, nil
}
func WriteCoverImage(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, 0644)
}

func main() {
	coverImagePath := "path/to/cover.jpg"
	// Reading cover image
	coverImage, err := ReadCoverImage(coverImagePath)
	if err != nil {
		fmt.Printf("Failed to read cover image: %s\n", err)
		return
	}
	// Assuming a book needs its cover image updated
	if err := WriteCoverImage(coverImagePath, coverImage); err != nil {
		fmt.Printf("Failed to write cover image: %s\n", err)
	}
}