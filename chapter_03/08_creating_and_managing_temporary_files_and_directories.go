package main

import (
	"fmt"
	"os"
)

func CreateTempFile(prefix string) (*os.File, error) {
	tempFile, err := os.CreateTemp("", prefix)
	if err != nil {
		return nil, err
	}
	// TempFile creates the file with os.O_RDWR|os.O_CREATE|os.O_EXCL mode
	return tempFile, nil
}
func CreateTempDir(prefix string) (string, error) {
	tempDir, err := os.MkdirTemp("", prefix)
	if err != nil {
		return "", err
	}
	return tempDir, nil
}
func ProcessAndCleanupTempFile(tempFile *os.File) {
	// Example processing on tempFile
	// ...
	// Cleanup
	defer os.Remove(tempFile.Name())
}
func ProcessAndCleanupTempDir(tempDir string) {
	// Example processing using tempDir
	// ...
	// Cleanup
	defer os.RemoveAll(tempDir)
}
func main() {
	tempFile, err := CreateTempFile("librago")
	if err != nil {
		fmt.Printf("Failed to create a temporary file: %s\n", err)
		return
	}
	ProcessAndCleanupTempFile(tempFile)
	tempDir, err := CreateTempDir("librago")
	if err != nil {
		fmt.Printf("Failed to create a temporary directory: %s\n", err)
		return
	}
	ProcessAndCleanupTempDir(tempDir)
}