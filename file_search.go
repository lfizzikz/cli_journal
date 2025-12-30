package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type SearchOptions struct {
	Date  string
	From  string
	To    string
	Tags  []string
	Query []string
}

func FilesToSearch(opts SearchOptions) ([]string, error) {
	var filestoSearch []string
	var dateSearch time.Time
	var toSearch time.Time
	var fromSearch time.Time
	var hasDate, hasTo, hasFrom bool
	files, err := os.ReadDir(VaultPath)
	if err != nil {
		return []string{}, err
	}
	if opts.Date != "" {
		dateSearch, err = time.Parse("2006-01-02", opts.Date)
		if err != nil {
			fmt.Println("date parse failed:", err)
		}
		hasDate = true
	}
	if opts.To != "" {
		toSearch, err = time.Parse("2006-01-02", opts.To)
		if err != nil {
			fmt.Println("to parse failed:", err)
		}
		hasTo = true
	}
	if opts.From != "" {
		fromSearch, err = time.Parse("2006-01-02", opts.From)
		if err != nil {
			fmt.Println("from parse failed:", err)
		}
		hasFrom = true
	}
	for _, file := range files {
		filename := file.Name()
		if strings.HasPrefix(filename, ".") || !strings.HasSuffix(filename, ".md") {
			continue
		}
		basename := strings.TrimSuffix(filename, ".md")
		fileDate, err := time.Parse("2006-01-02", basename)
		if err != nil {
			continue
		}
		switch {
		case hasDate:
			if fileDate.Equal(dateSearch) {
				filestoSearch = append(filestoSearch, filename)
			}
		case hasTo && hasFrom:
			if (fileDate.Equal(fromSearch) || fileDate.After(fromSearch)) &&
				(fileDate.Equal(toSearch) || fileDate.Before(toSearch)) {
				filestoSearch = append(filestoSearch, filename)
			}
		case hasTo:
			if fileDate.Before(toSearch) {
				filestoSearch = append(filestoSearch, filename)
			}
		case hasFrom:
			if fileDate.After(fromSearch) {
				filestoSearch = append(filestoSearch, filename)
			}
		}
	}
	return filestoSearch, nil
}
