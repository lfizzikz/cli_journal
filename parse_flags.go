package main

import (
	"flag"
	"strings"
)

func ParseSearchFlags(args []string) (SearchOptions, error) {
	fs := flag.NewFlagSet("search", flag.ContinueOnError)

	var opts SearchOptions
	var tagsList string
	var queryList string

	fs.StringVar(&opts.Date, "date", "", "search a specified date (YYYY-MM-DD)")
	fs.StringVar(&opts.From, "from", "", "search from date (YYYY-MM-DD)")
	fs.StringVar(&opts.To, "to", "", "search up to date (YYYY-MM-DD)")
	fs.StringVar(&tagsList, "tags", "", "search tags, comma-seperated (work, random)")
	fs.StringVar(&queryList, "query", "", "search words, comma-seperated (together, seperate)")

	if err := fs.Parse(args); err != nil {
		return SearchOptions{}, err
	}
	if tagsList != "" {
		opts.Tags = strings.Split(tagsList, ",")
	}
	if queryList != "" {
		opts.Query = strings.Split(queryList, ",")
	}

	return opts, nil
}

func ParseAddFlags(args []string) (tag []string, body string, err error) {
	fs := flag.NewFlagSet("add", flag.ContinueOnError)
	tagCSV := ""

	fs.StringVar(&tagCSV, "tag", "", "Will append #{tag} to end of entry")

	if err := fs.Parse(args); err != nil {
		return []string{}, "", err
	}

	if tagCSV != "" {
		rawTags := strings.Split(tagCSV, ",")
		for _, t := range rawTags {
			trimmed := strings.TrimSpace(t)
			if trimmed != "" {
				tag = append(tag, trimmed)
			}
		}
	}
	body = strings.Join(fs.Args(), " ")
	return tag, body, nil
}
