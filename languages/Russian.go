package languages

import (
	"regexp"
)

type Russian struct {
	Language
}


var RussianAbbreviations = []string{
	"y.e",
	"y",
	"а",
	"авт",
	"адм.-терр",
	"акад",
	"в",
	"вв",
	"вкз",
	"вост.-европ",
	"г",
	"гг",
	"гос",
	"гр",
	"д",
	"деп",
	"дисс",
	"дол",
	"долл",
	"ежедн",
	"ж",
	"жен",
	"з",
	"зап.-европ",
	"зап",
	"заруб",
	"и",
	"ин",
	"иностр",
	"инст",
	"к",
	"канд",
	"кв",
	"кг",
	"куб",
	"л.h",
	"л.н",
	"л",
	"м",
	"мин",
	"моск",
	"муж",
	"н",
	"нед",
	"о",
	"п",
	"пгт",
	"пер",
	"пп",
	"пр",
	"просп",
	"проф",
	"р",
	"руб",
	"с",
	"сек",
	"см",
	"спб",
	"стр",
	"т",
	"тел",
	"тов",
	"тт",
	"тыс",
	"у.е",
	"у",
	"ул",
	"ф",
	"ч",
}


type RuWordContinuityHelper struct {
}


func (helper *RuWordContinuityHelper) ContinueInNextWord(textAfterBoundary string) bool {
	return regexp.MustCompile(`^[0-9a-zа-я]`).MatchString(textAfterBoundary)
}


func (helper *RuWordContinuityHelper) GetLastWord(text string) string {
	words := regexp.MustCompile(`[\s\.]+`).Split(text, -1)
	return words[len(words)-1]
}


func NewRussian() *Russian {
	language := NewLanguage()
	language.Abbreviations = NewSetFromArray(RussianAbbreviations)
	language.IsPunctuationBetweenQuotes = false
	language.WordContinuityHelper = &RuWordContinuityHelper{}
	return &Russian{
		Language: *language,
	}
}
