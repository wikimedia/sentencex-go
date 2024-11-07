package languages

import (
	"regexp"
	"strings"
)

type Finnish struct {
	Language
}

var FI_MONTHS = []string{
	"tammikuu",
	"helmikuu",
	"maaliskuu",
	"huhtikuu",
	"toukokuu",
	"kesäkuu",
	"heinäkuu",
	"elokuu",
	"syyskuu",
	"lokakuu",
	"marraskuu",
	"joulukuu",
}

var FinnishAbbreviations = []string{
	"A",
	"B",
	"C",
	"D",
	"E",
	"F",
	"G",
	"H",
	"I",
	"J",
	"K",
	"L",
	"M",
	"N",
	"O",
	"P",
	"Q",
	"R",
	"S",
	"T",
	"U",
	"V",
	"W",
	"X",
	"Y",
	"Z",
	"Å",
	"Ä",
	"Ö",
	// List of titles.
	// These are often followed by upper-case names, but do not indicate sentence breaks
	"alik",
	"alil",
	"amir",
	"apul",
	"apul.prof",
	"arkkit",
	"ass",
	"assist",
	"dipl",
	"dipl.arkkit",
	"dipl.ekon",
	"dipl.ins",
	"dipl.kielenk",
	"dipl.kirjeenv",
	"dipl.kosm",
	"dipl.urk",
	"dos",
	"Dr",
	"erikoiseläinl",
	"erikoishammasl",
	"erikoisl",
	"erikoist",
	"ev.luutn",
	"evp",
	"fil",
	"ft",
	"hallinton",
	"hallintot",
	"hammaslääket",
	"jatk",
	"jääk",
	"kansaned",
	"kapt",
	"kapt.luutn",
	"kenr",
	"kenr.luutn",
	"kenr.maj",
	"kers",
	"kirjeenv",
	"kom",
	"kom.kapt",
	"komm",
	"konst",
	"korpr",
	"luutn",
	"maist",
	"maj",
	"Mr",
	"Mrs",
	"Ms",
	"M.Sc",
	"neuv",
	"nimim",
	"Ph.D",
	"prof",
	"puh.joht",
	"pääll",
	"res",
	"san",
	"siht",
	"suom",
	"sähköp",
	"säv",
	"toht",
	"toim",
	"toim.apul",
	"toim.joht",
	"toim.siht",
	"tuom",
	"ups",
	"vänr",
	"vääp",
	"ye.ups",
	"ylik",
	"ylil",
	"ylim",
	"ylimatr",
	"yliop",
	"yliopp",
	"ylip",
	"yliv",
	// misc - odd period-ending items that NEVER indicate breaks (p.m. does NOT fall
	// into this category - it sometimes ends a sentence)
	"e.g",
	"ent",
	"esim",
	"huom",
	"i.e",
	"ilm",
	"l",
	"mm",
	"myöh",
	"nk",
	"nyk",
	"par",
	"po",
	"t",
	"v",
}

type FiWordContinuityHelper struct {
}


func (helper *FiWordContinuityHelper) ContinueInNextWord(textAfterBoundary string) bool {
	re := regexp.MustCompile(`^\W*[0-9a-z]`)
	if re.MatchString(textAfterBoundary) {
		return true
	}

	words := regexp.MustCompile(`\s+`).Split(strings.TrimSpace(textAfterBoundary), -1)

	if len(words) == 0 {
		return false
	}

	nextWord := words[0]
	nextWord = strings.Trim(nextWord, "?!.")

	if len(nextWord) == 0 {
		return false
	}

	for _, month := range FI_MONTHS {
		if nextWord == month || (strings.ToUpper(string(nextWord[0]))+nextWord[1:] == month) {
			return true
		}
	}

	return false
}


func (helper *FiWordContinuityHelper) GetLastWord(text string) string {
	words := regexp.MustCompile(`[\s\.]+`).Split(text, -1)
	return words[len(words)-1]
}

func NewFinnish() *Finnish {
	language := NewLanguage()
	language.Abbreviations = NewSetFromArray(append(FinnishAbbreviations, EnAbbreviations...))
	language.IsPunctuationBetweenQuotes = true
	language.WordContinuityHelper = &FiWordContinuityHelper{}
	return &Finnish{
		Language: *language,
	}
}
