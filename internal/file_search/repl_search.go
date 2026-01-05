package filesearch

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ListFilesAndSearch(files []SearchResult) (string, error) {
	scanner := bufio.NewScanner(os.Stdin)

	for i, f := range files {
		fmt.Printf("%d) %s -%s\n", i+1, f.File, f.FirstSentence)
	}

	for {
		fmt.Println("")
		fmt.Print("Input number to open (q to quit) > ")

		ok := scanner.Scan()
		if !ok {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		numInput, _ := strconv.Atoi(input)
		if input == "" {
			continue
		}
		if input == "q" {
			os.Exit(1)
		}

		if numInput > len(files) {
			fmt.Println("Enter a valid number")
			continue
		}
		fileNumber, err := strconv.Atoi(input)
		if err != nil {
			return "", err
		}
		return files[fileNumber-1].File, nil
	}
	return "", nil
}
