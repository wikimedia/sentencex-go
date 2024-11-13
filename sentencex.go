package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/wikimedia/sentencex-go/languages"
)

var consecutiveNewlineRegex = regexp.MustCompile(`[\r\n]{2}`)

// Modified version of Go's builtin bufio.ScanLines to return strings separated by
// two newlines (instead of one).
// https://github.com/golang/go/blob/master/src/bufio/scan.go#L344-L364
func ScanTwoConsecutiveNewlines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if loc := consecutiveNewlineRegex.FindIndex(data); loc != nil && loc[0] >= 0 {
		return loc[1], data[0 : loc[0]+2], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	// Request more data.
	return 0, nil, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: sentencex <language> [files...]")
		return
	}

	language := os.Args[1]
	files := os.Args[2:]

	var scanner *bufio.Scanner

	if len(files) == 0 {
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", file, err)
				continue
			}
			defer f.Close()
			scanner = bufio.NewScanner(f)
			scanner.Split(ScanTwoConsecutiveNewlines)
			processFile(scanner, language)
		}
		return
	}

	processFile(scanner, language)
}

func processFile(scanner *bufio.Scanner, language string) {
	for scanner.Scan() {
		text := scanner.Text()
		sentences := segment(language, text)
		for _, sentence := range sentences {
			fmt.Printf("> %s\n", sentence)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}

func segment(language, text string) []string {
	factory := languages.LanguageFactory{}
	lang := factory.CreateLanguage(language)
	return lang.Segment(text)
}
