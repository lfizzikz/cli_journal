package filesearch

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ListFilesAndSearch(files []string) (string, error) {
	scanner := bufio.NewScanner(os.Stdin)

	for i, f := range files {
		fmt.Printf("%d) %s\n", i+1, f)
	}

	for {
		fmt.Println("")
		fmt.Print("Select a number to open (or q to quit) > ")

		ok := scanner.Scan()
		if !ok {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}
		if input == "q" {
			fmt.Println("Exiting")
			break
		}

		fileNumber, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter only numbers shown")
			return "", err
		}
		return files[fileNumber-1], nil
	}
	return "", nil
}
