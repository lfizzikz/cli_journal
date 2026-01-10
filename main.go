package main

import (
	"fmt"
	"odn/internal/config"
	filesearch "odn/internal/file_search"
	parseflags "odn/internal/parse_flags"
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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected command: search | add | open")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "search":
		opts, err := parseflags.ParseSearchFlags(os.Args[2:])
		if err != nil {
			fmt.Println("error after parse search:", err)
			os.Exit(1)
		}

		files, err := filesearch.FilesToSearch(opts)
		if err != nil {
			fmt.Println("error after finding files to search:", err)
			os.Exit(1)
		}
		if len(files) == 0 {
			fmt.Println("no files found.")
		} else {
			foundFiles, err := filesearch.SearchInFile(files, opts)
			if err != nil {
				fmt.Println("error on search:", err)
				os.Exit(1)
			}
			fileToOpen, err := filesearch.ListFilesAndSearch(foundFiles)
			if err != nil {
				fmt.Println("error on listing files:", err)
			}
			openInObsidian(fileToOpen)
		}
	case "add":
		tag, body, err := parseflags.ParseAddFlags(os.Args[2:])
		if err != nil {
			fmt.Println("add parse error:", err)
			os.Exit(1)
		}
		fTime, fDate := getDateTime()
		newFile := createNewFileStruct(fTime, fDate, body, tag)
		writeToFile(newFile)
	case "open":
		file, err := parseflags.ParseOpenFlags(os.Args[2:])
		if err != nil {
			fmt.Println("open parse error", err)
			os.Exit(1)
		}
		if err = openInObsidian(file); err != nil {
			fmt.Println("open obsidian error:", err)
			os.Exit(1)
		}
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
	fullPath := config.VaultPath + date + extension

	newFile := FileInfo{
		content:   entry,
		title:     date,
		fullPath:  fullPath,
		entryTime: time,
		tags:      tags,
	}
	return newFile
}
