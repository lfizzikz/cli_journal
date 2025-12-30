package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type FileInfo struct {
	content   string
	title     string
	fullPath  string
	entryTime string
	tags      []string
}

const VaultPath = "/Users/trevornance/Documents/My Vault/Daily Writing/"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected command: search | add")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "search":
		opts, err := ParseSearchFlags(os.Args[2:])
		if err != nil {
			fmt.Println("error after parse search:", err)
			os.Exit(1)
		}

		files, err := FilesToSearch(opts)
		if err != nil {
			fmt.Println("error after finding files to search:", err)
			os.Exit(1)
		}
		if len(files) == 0 {
			fmt.Println("no files found.")
		} else {
			fmt.Println(files)
		}
	case "add":
		tag, body, err := ParseAddFlags(os.Args[2:])
		if err != nil {
			fmt.Println("add parse error:", err)
			os.Exit(1)
		}
		fTime, fDate := getDateTime()
		newFile := createNewFileStruct(fTime, fDate, body, tag)
		writeToFile(newFile)
	}
}

func writeToFile(f FileInfo) {
	file, err := os.OpenFile(
		f.fullPath,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		fmt.Printf("Error: %s\nFile: %s", err, f.title)
	}
	defer file.Close()
	if len(f.tags) == 0 {
		contentToWrite := fmt.Sprintf("- [%s] %s\n\n", f.entryTime, f.content)
		file.WriteString(contentToWrite)
	} else {
		tagText := "#" + strings.Join(f.tags, " #")
		contentToWrite := fmt.Sprintf("- [%s] %s %s\n\n", f.entryTime, f.content, tagText)
		file.WriteString(contentToWrite)
	}
}

func getDateTime() (fTime, fDate string) {
	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04")
	return formattedTime, formattedDate
}

func createNewFileStruct(time, date, entry string, tags []string) FileInfo {
	extension := ".md"
	fullPath := VaultPath + date + extension

	newFile := FileInfo{
		content:   entry,
		title:     date,
		fullPath:  fullPath,
		entryTime: time,
		tags:      tags,
	}
	return newFile
}
