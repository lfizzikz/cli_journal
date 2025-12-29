package main

import (
	"fmt"
	"os"
)

type SearchOptions struct {
	Date  string
	From  string
	To    string
	Tags  []string
	Query []string
}

func FilesToSearch(opts SearchOptions) ([]string, error) {
	files, err := os.ReadDir(VaultPath)
	if err != nil {
		return []string{}, err
	}
	for _, file := range files {
		fmt.Println(file)
	}
	return []string{}, nil
	// 1. Read directory listing (filenames only)
	// 2. Extract date from each filename
	// 3. Convert extracted date → real date type
	// 4. Compare against from/to dates
	// 5. Only then open the matching files (if needed)
}
