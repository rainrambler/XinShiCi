package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

// https://golang.cafe/blog/how-to-list-files-in-a-directory-in-go.html
// fileext: ".txt"
func FindFilesInDir(dir, fileext string) []string {
	allres := []string{}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(info.Name()) != fileext {
			return nil
		}

		allres = append(allres, path)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return allres
}

// https://golang.cafe/blog/how-to-list-files-in-a-directory-in-go.html
func FindAllFilesInDir(dir string) []string {
	allres := []string{}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		allres = append(allres, path)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return allres
}

// https://golangr.com/rename-file/
func renameFile(src, dst string) {
	// rename file
	os.Rename(src, dst)
}

func checkFileExists(filePath string) bool {
	_, error := os.Stat(filePath)
	//return !os.IsNotExist(err)
	return !errors.Is(error, os.ErrNotExist)
}
