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
}

var VaultPath = "/Users/trevornance/Documents/My Vault/Daily Writing/"

func main() {

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

	file.WriteString("-[" + f.entryTime + "] " + f.content + "\n\n")
}

func getDateTime() (fTime, fDate string) {
	currentDateTime := time.Now()
	formattedDate := currentDateTime.Format("2006-01-02")
	formattedTime := currentDateTime.Format("15:04")
	return formattedTime, formattedDate
}

func createNewFileStruct(time, date string) FileInfo {
	extension := ".md"
	entry := os.Args[1:]
	entryToSave := strings.Join(entry, " ")
	fullPath := VaultPath + date + extension

	newFile := FileInfo{
		content:   entryToSave,
		title:     date,
		fullPath:  fullPath,
		entryTime: time,
	}
	return newFile
}
