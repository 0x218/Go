package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func copyFolder(srcPath, destPath string) error {
	if err := os.MkdirAll(destPath, 0755); err != nil {
		return err //destination path doesn't exists
	}

	err := filepath.Walk(srcPath, func(sourcePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err //unable to read source path
		}

		relativePath, err := filepath.Rel(srcPath, sourcePath)
		if err != nil {
			return err //unable to locate relative path
		}

		//create desitnation path
		destinationPath := filepath.Join(destPath, relativePath)

		if info.IsDir() {
			return os.MkdirAll(destinationPath, info.Mode())
		}

		emptyFile, err := os.Create(destinationPath)
		if err != nil {
			return err
		}
		defer emptyFile.Close()

		return nil
	})
	return err
}

func main() {
	currentDir, err := os.Getwd()

	if err != nil {
		fmt.Println("Error getting current directory", err)
		os.Exit(1)
	}

	destDir := filepath.Join(currentDir, "StructWithEmptyFileContent")

	err = copyFolder(currentDir, destDir)
	if err != nil {
		fmt.Println("Error copying folder structure", err)
		os.Exit(1)
	}
	fmt.Println("Folder structure copied to ", destDir)
}
