package main

import (
	"os"
	"strings"
	"time"
)

type FileInfo struct {
	content string
	title   string
}

var VaultPath = "/Users/trevornance/Documents/My Vault/Daily Writing/"

func main() {
	extension := ".md"
	entry := os.Args[1:]
	currentDate := time.Now()
	formattedDate := currentDate.Format("2006-01-02")
	entryToSave := strings.Join(entry, " ")
	fullPath := VaultPath + formattedDate + extension
	newFile := FileInfo{
		content: entryToSave,
		title:   fullPath,
	}
	file, err := os.Create(newFile.title)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.WriteString(newFile.content)
}
