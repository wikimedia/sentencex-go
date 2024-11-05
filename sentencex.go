package sentencex

import (
	"bufio"
	"fmt"
	"os"

	"github.com/wikimedia/sentencex-go/languages"
)

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
			fmt.Printf("|%s|\n", sentence)
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
