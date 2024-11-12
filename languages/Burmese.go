package languages

import (
	"fmt"
	"regexp"
	"strings"
)

type Burmese struct {
	Language
}

func NewBurmese() *Burmese {
	language := NewLanguage()
	global_patterns := GLOBAL_SENTENCE_TERMINATORS
	global_patterns = append(global_patterns, "၏")
	language.SentenceBreakRegex = regexp.MustCompile(fmt.Sprintf("[%s]+", strings.Join(global_patterns, "")))
	return &Burmese{
		Language: *language,
	}
}
