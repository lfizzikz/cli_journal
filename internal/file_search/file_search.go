package filesearch

import (
	"bufio"
	"fmt"
	"odn/internal/config"
	"os"
	"strings"
	"time"
)

type SearchOptions struct {
	Date     string
	Year     string
	Month    string
	From     string
	To       string
	Tags     []string
	Query    []string
	BodyText string
}

type SearchResult struct {
	File          string
	FirstSentence string
}

func FilesToSearch(opts SearchOptions) ([]string, error) {
	var filestoSearch []string
	var dateSearch time.Time
	var yearSearch time.Time
	var monthSearch time.Time
	var toSearch time.Time
	var fromSearch time.Time
	var hasDate, hasTo, hasFrom, hasYear, hasMonth bool
	files, err := os.ReadDir(config.VaultPath)
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
	if opts.Year != "" {
		yearSearch, err = time.Parse("2006", opts.Year)
		if err != nil {
			fmt.Println("year parse failed:", err)
		}
		hasYear = true
	}
	if opts.Month != "" {
		monthSearch, err = time.Parse("January", opts.Month)
		if err != nil {
			monthSearch, err = time.Parse("01", opts.Month)
		}
		if err != nil {
			monthSearch, err = time.Parse("1", opts.Month)
		}
		if err != nil {
			fmt.Println("month parse failed:", err)
		} else {
			hasMonth = true
		}
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
	noDateFilters := opts.Date == "" && opts.From == "" && opts.Year == "" && opts.To == "" && opts.Month == ""
	for _, file := range files {
		filename := file.Name()
		if strings.HasPrefix(filename, ".") || !strings.HasSuffix(filename, ".md") {
			continue
		}
		basename := strings.TrimSuffix(filename, ".md")
		if noDateFilters {
			filestoSearch = append(filestoSearch, filename)
		}
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
		case hasYear && hasMonth:
			if (fileDate.Month() == monthSearch.Month()) &&
				(fileDate.Year() == yearSearch.Year()) {
				filestoSearch = append(filestoSearch, filename)
			}
		case hasYear:
			if fileDate.Year() == yearSearch.Year() {
				filestoSearch = append(filestoSearch, filename)
			}
		case hasMonth:
			if fileDate.Month() == monthSearch.Month() {
				filestoSearch = append(filestoSearch, filename)
			}
		}
	}
	return filestoSearch, nil
}

func SearchInFile(files []string, opts SearchOptions) ([]SearchResult, error) {
	searchFound := []SearchResult{}
	for _, file := range files {
		contains, firstSentence, err := FileContainsAll(file, opts)
		if err != nil {
			return []SearchResult{}, err
		}
		if contains {
			searchFound = append(searchFound, SearchResult{
				File:          file,
				FirstSentence: firstSentence,
			})
		}
	}
	return searchFound, nil
}

func FileContainsAll(file string, opts SearchOptions) (bool, string, error) {
	found := make(map[string]bool)
	file = config.VaultPath + file
	f, err := os.Open(file)
	if err != nil {
		return false, "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	firstSentence := getFirstSentence(scanner)

	for scanner.Scan() {
		line := scanner.Text()

		for _, q := range opts.Query {
			if strings.Contains(line, q) {
				found["q:"+q] = true
			}
		}

		for _, t := range opts.Tags {
			tags := "#" + t
			if strings.Contains(line, tags) {
				found["t:"+t] = true
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return false, "", err
	}

	for _, q := range opts.Query {
		if !found["q:"+q] {
			return false, "", err
		}
	}

	for _, t := range opts.Tags {
		if !found["t:"+t] {
			return false, "", err
		}
	}
	return true, firstSentence, nil
}

func getFirstSentence(scanner *bufio.Scanner) (sentence string) {
	for scanner.Scan() {
		line := scanner.Text()
		for i, r := range line {
			if r == '.' || r == '!' || r == '?' {
				return strings.TrimSpace(line[:i+1])
			}
		}
	}
	return "error getting sentence"
}
