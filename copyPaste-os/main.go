package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	srcDir := "./common"
	newName := "NewPage.jsx"

	destDir := "." // Default destination directory

	// Forming destination file path
	destPath := filepath.Join(destDir, newName)

	// Copying file
	err := copyFile(filepath.Join(srcDir, "Page.jsx"), destPath)
	if err != nil {
		log.Fatalf("Error copying file: %s", err)
	}

	log.Printf("File copied from '%s' to '%s'", srcDir, destPath)
}
