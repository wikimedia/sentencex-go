package languages

type ILanguage interface {
	Segment(text string) []string
	GetSkippableRanges(text string) [][2]int
	ContinueInNextWord(textAfterBoundary string) bool
	FindBoundary(text string, start int, end int) int
	GetLastWord(text string) string
	IsAbbreviation(head, tail, separator string) bool
	IsExclamationWord(head, tail string) bool
}
