package main

import (
	"flag"
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
	tags      string
}

const VaultPath = "/Users/trevornance/Documents/My Vault/Daily Writing/"

func main() {
	searchFlag := flag.String("search", "", "search file contents. {date}, {from}, {to}, {[tags]}, {[query]}")
	fDate, fTime := getDateTime()
	newFile := createNewFileStruct(fDate, fTime)
	writeToFile(newFile)
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
	contentToWrite := fmt.Sprintf("- [%s] %s #%s\n\n", f.entryTime, f.content, f.tags)
	file.WriteString(contentToWrite)
}

func getDateTime() (fTime, fDate string) {
	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04")
	return formattedTime, formattedDate
}

func createNewFileStruct(time, date string) FileInfo {
	tagsFlag := flag.String("tag", "", "Will append #{tag} to end of entry")
	flag.Parse()
	tags := *tagsFlag
	extension := ".md"
	entry := strings.Join(flag.Args(), " ")
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
