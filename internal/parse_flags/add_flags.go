package parseflags

import (
	"flag"
	"strings"
)

func ParseAddFlags(args []string) (tag []string, body string, needsInput bool, err error) {
	fs := flag.NewFlagSet("add", flag.ContinueOnError)
	tagCSV := ""

	fs.StringVar(&tagCSV, "tag", "", "Will append #{tag} to end of entry")

	if err := fs.Parse(args); err != nil {
		return []string{}, "", false, err
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
	remaining := fs.Args()
	if len(remaining) > 0 {
		body = strings.Join(remaining, " ")
		return tag, body, false, nil
	}
	return tag, "", true, nil
}
