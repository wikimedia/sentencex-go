package languages

import (
	"fmt"
	"regexp"
	"strings"
)

type IWordContinuityHelper interface {
	ContinueInNextWord(textAfterBoundary string) bool
	GetLastWord(text string) string
}

type Language struct {
	Language                   string
	QuotePairs                 map[string]string
	GlobalSentenceBoundaryReg  *regexp.Regexp
	QuotesRegex                *regexp.Regexp
	ParensRegex                *regexp.Regexp
	EmailRegex                 *regexp.Regexp
	NumberedReferenceRegex     *regexp.Regexp
	SpaceAfterSeperator        *regexp.Regexp
	SentenceBreakRegex         *regexp.Regexp
	Abbreviations              map[string]struct{}
	AbbreviationChar           string
	ExclamationWords           map[string]struct{}
	IsPunctuationBetweenQuotes bool
	WordContinuityHelper       IWordContinuityHelper
}

func NewLanguage() *Language {
	exclamationWords := map[string]struct{}{
		"!Xũ": {}, "!Kung": {}, "ǃʼOǃKung": {}, "!Xuun": {}, "!Kung-Ekoka": {}, "ǃHu": {}, "ǃKhung": {}, "ǃKu": {}, "ǃung": {}, "ǃXo": {}, "ǃXû": {}, "ǃXung": {}, "ǃXũ": {}, "!Xun": {}, "Yahoo!": {}, "Y!J": {}, "Yum!": {},
	}

	quotePairs := map[string]string{
		`"`:  `"`,
		" '": "'", // # Need a space before ' to avoid capturing don't , l'Avv etc
		"«":  "»",
		"‘":  "’",
		"‚":  "‚",
		"“":  "”",
		"‛":  "‛",
		"„":  "“",
		"‟":  "‟",
		"‹":  "›",
		"《":  "》",
		"「":  "」",
	}

	globalSentenceBoundaryReg := regexp.MustCompile(fmt.Sprintf("[%s]+", strings.Join(GLOBAL_SENTENCE_TERMINATORS, "")))
	if globalSentenceBoundaryReg == nil {
		panic("globalSentenceBoundaryReg is nil")
	}
	//quotes_regx_str = r"|".join([f"{left}(\n|.)*?{right}" for left, right in quote_pairs.items()])
	quotesRegxStr := strings.Join(func() []string {
		var pairs []string
		for left, right := range quotePairs {
			pairs = append(pairs, left+"(\n|.)*?"+right)
		}
		return pairs
	}(), "|")
	fmt.Println(quotesRegxStr)
	quotesRegex, err := regexp.Compile(quotesRegxStr)
	if err != nil {
		fmt.Println("Error compiling regular expression:", err)

	}
	parensRegex := regexp.MustCompile(`\([^)]+\)`)
	emailRegex := regexp.MustCompile(`[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Z|a-z]{2,7}`)
	numberedReferenceRegex := regexp.MustCompile(`^(\[\d+])+`)
	spaceAfterSeperator := regexp.MustCompile(`^\s+`)

	return &Language{
		Language:                   "base",
		QuotePairs:                 quotePairs,
		GlobalSentenceBoundaryReg:  globalSentenceBoundaryReg,
		QuotesRegex:                quotesRegex,
		ParensRegex:                parensRegex,
		EmailRegex:                 emailRegex,
		NumberedReferenceRegex:     numberedReferenceRegex,
		SentenceBreakRegex:         globalSentenceBoundaryReg,
		Abbreviations:              make(map[string]struct{}),
		AbbreviationChar:           ".",
		ExclamationWords:           exclamationWords,
		SpaceAfterSeperator:        spaceAfterSeperator,
		IsPunctuationBetweenQuotes: false,
	}
}

func (l *Language) IsAbbreviation(head, tail, separator string) bool {
	if l.AbbreviationChar != separator {
		return false
	}

	lastWord := l.GetLastWord(head)

	if len(lastWord) == 0 {
		return false
	}

	_, isAbbrev := l.Abbreviations[lastWord]
	_, isAbbrevLower := l.Abbreviations[strings.ToLower(lastWord)]
	_, isAbbrevUpper := l.Abbreviations[strings.ToUpper(lastWord)]

	return isAbbrev || isAbbrevLower || isAbbrevUpper
}

func (l *Language) IsExclamationWord(head, tail string) bool {
	lastWord := l.GetLastWord(head)
	_, exists := l.ExclamationWords[lastWord+"!"]
	return exists
}

func (l *Language) GetLastWord(text string) string {
	if l.WordContinuityHelper != nil {
		return l.WordContinuityHelper.GetLastWord(text)
	}
	words := regexp.MustCompile(`[\s\.]+`).Split(text, -1)
	return words[len(words)-1]
}

func (l *Language) FindBoundary(text string, start int, end int) int {
	head := text[:start]
	tail := text[start+1:]

	numberRefMatch := l.NumberedReferenceRegex.FindString(tail)
	if numberRefMatch != "" {
		return start + 1 + len(numberRefMatch)
	}

	if l.ContinueInNextWord(tail) {
		return -1
	}

	if l.IsAbbreviation(head, tail, text[start:end]) {
		return -1
	}
	if l.IsExclamationWord(head, tail) {
		return -1
	}

	spaceAfterSepMatch := l.SpaceAfterSeperator.FindString(tail)
	if spaceAfterSepMatch != "" {
		return start + 1 + len(spaceAfterSepMatch)
	}

	return end
}

func (l *Language) ContinueInNextWord(textAfterBoundary string) bool {
	if l.WordContinuityHelper != nil {
		return l.WordContinuityHelper.ContinueInNextWord(textAfterBoundary)
	}

	return regexp.MustCompile(`^[0-9a-z]`).MatchString(textAfterBoundary)
}

func (l *Language) GetSkippableRanges(text string) [][2]int {
	var skippableRanges [][2]int
	for _, match := range l.QuotesRegex.FindAllStringIndex(text, -1) {
		fmt.Println(match)
		skippableRanges = append(skippableRanges, [2]int{match[0], match[1]})
	}
	for _, match := range l.ParensRegex.FindAllStringIndex(text, -1) {
		skippableRanges = append(skippableRanges, [2]int{match[0], match[1]})
	}
	for _, match := range l.EmailRegex.FindAllStringIndex(text, -1) {
		skippableRanges = append(skippableRanges, [2]int{match[0], match[1]})
	}

	return skippableRanges
}

func (l *Language) Segment(text string) []string {
	var sentences []string
	paragraphs := regexp.MustCompile(`\n{2}`).Split(text, -1)

	for pindex, paragraph := range paragraphs {
		if pindex > 0 {
			sentences = append(sentences, "\n\n")
		}

		// Initialize a list to store the boundaries of sentences.
		boundaries := []int{0}

		if l.SentenceBreakRegex == nil {
			panic("SentenceBreakRegex is nil")
		}
		matches := l.SentenceBreakRegex.FindAllStringIndex(paragraph, -1)
		skippableRanges := l.GetSkippableRanges(paragraph)

		for _, match := range matches {
			boundary := l.FindBoundary(paragraph, match[0], match[1])
			// If boundary is 0, skip to the next match.
			if boundary == -1 {
				continue
			}

			// Check if the boundary is inside a skippable range (quote, parentheses, or email).
			inRange := false

			for _, rng := range skippableRanges {
				// Convert the rng to rune slice to avoid issues with multi-byte characters.

				rangeStart := len([]rune(paragraph[:rng[0]]))
				rangeEnd := len([]rune(paragraph[:rng[1]]))

				boundaryEnd := len([]rune(paragraph[:boundary]))

				if boundaryEnd > rangeStart && boundaryEnd < rangeEnd {
					if boundaryEnd+1 == rangeEnd && l.IsPunctuationBetweenQuotes {
						boundary = rng[1]
						inRange = false
					} else {
						inRange = true
					}

					break
				}
			}
			if inRange {
				continue
			}

			boundaries = append(boundaries, boundary)
		}
		// Add the last boundary to the list if it's not already there.
		if boundaries[len(boundaries)-1] != len(paragraph) {
			boundaries = append(boundaries, len(paragraph))
		}

		// fmt.Println(boundaries)
		for i := 0; i < len(boundaries)-1; i++ {
			start := boundaries[i]
			end := boundaries[i+1]

			// Ensure we don't go out of bounds
			if start >= len(paragraph) || end > len(paragraph) || start > end {
				continue
			}

			sentence := paragraph[start:end]
			sentences = append(sentences, sentence)
		}

	}

	return sentences
}

// NewSetFromArray creates a new set from a slice of any comparable type
func NewSetFromArray(arr []string) map[string]struct{} {
	// Initialize the set
	set := make(map[string]struct{})

	// Add each element to the set
	for _, item := range arr {
		set[item] = struct{}{}
	}

	return set
}
