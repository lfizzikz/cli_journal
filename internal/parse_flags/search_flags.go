package parseflags

import (
	"flag"
	filesearch "odn/internal/file_search"
	"strings"
)

func ParseSearchFlags(args []string) (filesearch.SearchOptions, error) {
	fs := flag.NewFlagSet("search", flag.ContinueOnError)

	var opts filesearch.SearchOptions
	var tagsList string
	var queryList string

	fs.StringVar(&opts.Date, "date", "", "search a specified date (YYYY-MM-DD)")
	fs.StringVar(&opts.Year, "year", "", "search a specified year (YYYY)")
	fs.StringVar(&opts.Month, "month", "", "search a specified month (MM)")
	fs.StringVar(&opts.From, "from", "", "search from date (YYYY-MM-DD)")
	fs.StringVar(&opts.To, "to", "", "search up to date (YYYY-MM-DD)")
	fs.StringVar(&tagsList, "tags", "", "search tags, comma-seperated (work, random)")
	fs.StringVar(&queryList, "query", "", "search words, comma-seperated (together, seperate)")

	if err := fs.Parse(args); err != nil {
		return filesearch.SearchOptions{}, err
	}
	if tagsList != "" {
		opts.Tags = strings.Split(tagsList, ",")
	}
	if queryList != "" {
		opts.Query = strings.Split(queryList, ",")
	}

	return opts, nil
}
