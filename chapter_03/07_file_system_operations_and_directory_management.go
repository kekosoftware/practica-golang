package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func OrganizeBooksByAuthor(libraryPath string, books []Book) error {
	for _, book := range books {
		authorDir := filepath.Join(libraryPath, sanitizeFileName(book.Author))
		if err := os.MkdirAll(authorDir, 0755); err != nil {
			return err
		}
		originalPath := filepath.Join(libraryPath, book.FileName)
		newPath := filepath.Join(authorDir, book.FileName)
		if err := os.Rename(originalPath, newPath); err != nil {
			return err
		}
	}
	return nil
}
func sanitizeFileName(name string) string {
	// Implement filename sanitization to remove/replace invalid characters
	// This is platform-dependent and left as an exercise
	return name
}
func CleanupEmptyDirectories(rootDir string) error {
	return filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			entries, err := os.ReadDir(path)
			if err != nil {
				return err
			}
			if len(entries) == 0 && path != rootDir {
				if err := os.Remove(path); err != nil {
					return err
				}
			}
		}
		return nil
	})
}
func main() {
	libraryPath := "/path/to/digital/library"
	// Assuming books is populated with the user's digital book collection
	if err := OrganizeBooksByAuthor(libraryPath, books); err != nil {
		fmt.Printf("Failed to organize books by author: %s\n", err)
	}
	if err := CleanupEmptyDirectories(libraryPath); err != nil {
		fmt.Printf("Failed to clean up empty directories: %s\n", err)
	}
}
