package languages

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

type Armenian struct {
	Language
}

func NewArmenian() *Armenian {
	language := NewLanguage()
	global_patterns := append(GLOBAL_SENTENCE_TERMINATORS, "։", "՜", ":")
	// remove "." from global patterns
	global_patterns = slices.DeleteFunc(global_patterns, func(cmp string) bool {
		return cmp == "."
	})
	language.SentenceBreakRegex = regexp.MustCompile(fmt.Sprintf("[%s]+", strings.Join(global_patterns, "")))
	return &Armenian{
		Language: *language,
	}
}
