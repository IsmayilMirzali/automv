package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	homedir, err              = os.UserHomeDir()
	documentsDir       string = filepath.Join(homedir, "docs")
	downloadsDir       string = filepath.Join(homedir, "dl")
	picturesDir        string = filepath.Join(homedir, "pics")
	documentExtensions        = []string{".pdf", ".doc", ".docx", ".odt", ".txt", ".rtf", ".tex"}
	pictureExtensions         = []string{".png", ".jpeg", ".jpg", ".gif"}
	moveCount          int32  = 0
)

func main() {
	files, err := ioutil.ReadDir(downloadsDir)
	if err != nil {
		log.Fatalln(err)
	}

	pictures := findMatchingFiles(files, pictureExtensions)
	documents := findMatchingFiles(files, documentExtensions)
	if pictures != nil {
		moveFiles(pictures, picturesDir)
	}
	if documents != nil {
		moveFiles(documents, documentsDir)
	}
	log.Printf("%d files moved.\n", moveCount)
}

func moveFiles(files []string, destination string) {
	for _, file := range files {
		if file != "" {
			os.Rename(filepath.Join(downloadsDir, file), filepath.Join(destination, file))
			moveCount++
		}
	}
}

/*
	Find matching files based on a list of files and file extensions to match
*/
func findMatchingFiles(files []os.FileInfo, extensions []string) []string {
	if len(files) < 1 || len(extensions) < 1 {
		return nil
	}
	matches := make([]string, len(files))
	for _, file := range files {
		for _, exten := range extensions {
			if strings.Contains(exten, filepath.Ext(file.Name())) {
				matches = append(matches, file.Name())
			}
		}
	}
	return matches
}
