package parseflags

import (
	"flag"
	"strings"
	"time"
)

func ParseOpenFlags(args []string) (file string, err error) {
	fs := flag.NewFlagSet("open", flag.ContinueOnError)

	if err := fs.Parse(args); err != nil {
		return "", err
	}
	if file == "" {
		rest := fs.Args()
		if len(rest) > 0 {
			file = rest[0]
		}
	}
	file = strings.TrimSpace(file)
	if file == "today" {
		file = time.Now().Format("2006-01-02")
	}
	if file != "" && !strings.HasSuffix(file, ".md") {
		file += ".md"
	}
	return file, nil
}
