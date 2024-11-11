package languages

import (
	"regexp"
	"strings"
)

type Deutsch struct {
	Language
}

var DE_MONTHS = []string{
	"Januar",
	"Februar",
	"März",
	"April",
	"Mai",
	"Juni",
	"Juli",
	"August",
	"September",
	"Oktober",
	"November",
	"Dezember",
}

var DeutschAbbreviations = []string{
	"ä",
	"Ä",
	"adj",
	"adm",
	"adv",
	"ao.univ.prof",
	"art",
	"ass.prof",
	"ass",
	"asst",
	"b.a",
	"b.s",
	"bart",
	"bldg",
	"brig",
	"bros",
	"bse",
	"buchst",
	"bzgl",
	"bzw",
	"c.-à-d",
	"ca",
	"capt",
	"chr",
	"cmdr",
	"co",
	"col",
	"comdr",
	"con",
	"corp",
	"cpl",
	"d.h",
	"d.j",
	"dergl",
	"dgl",
	"di",
	"dipl.-ing",
	"dkr",
	"dr ",
	"ens",
	"etc",
	"ev ",
	"evtl",
	"ff",
	"g.g.a",
	"g.u",
	"gen",
	"ggf",
	"gov",
	"hon.prof",
	"hon",
	"hosp",
	"i.f",
	"i.h.v",
	"ii",
	"iii",
	"insp",
	"iv",
	"ix",
	"jun",
	"k.o",
	"kath",
	"lfd",
	"lt",
	"ltd",
	"m.e",
	"mag",
	"maj",
	"med",
	"messrs",
	"mio",
	"mlle",
	"mm",
	"mme",
	"mr",
	"mrd",
	"mrs",
	"ms",
	"msgr",
	"mwst",
	"no",
	"nos",
	"nr",
	"o.ä",
	"o.univ.-prof",
	"op",
	"ord",
	"pfc",
	"ph",
	"pp",
	"prof",
	"projektass",
	"pvt",
	"rep",
	"reps",
	"res",
	"rev",
	"rt",
	"s",
	"s.p.a",
	"sa",
	"sen",
	"sens",
	"sfc",
	"sgt",
	"sog",
	"sogen",
	"spp",
	"sr",
	"st",
	"std",
	"str  ",
	"stud.ass",
	"supt",
	"surg",
	"T",
	"u.a  ",
	"u.ä",
	"u.e",
	"u.s.w",
	"u.u",
	"univ.-doz",
	"univ.-prof",
	"univ.ass",
	"usf",
	"usw",
	"v",
	"vgl",
	"vi",
	"vii",
	"viii",
	"vs",
	"x",
	"xi",
	"xii",
	"xiii",
	"xiv",
	"xix",
	"xv",
	"xvi",
	"xvii",
	"xviii",
	"xx",
	"z.b",
	"z.t",
	"z.z",
	"z.zt",
	"zt",
	"zzt",
}

type DeWordContinuityHelper struct {
}

func (helper *DeWordContinuityHelper) ContinueInNextWord(textAfterBoundary string) bool {
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

	for _, month := range DE_MONTHS {
		if nextWord == month || (strings.ToUpper(string(nextWord[0]))+nextWord[1:] == month) {
			return true
		}
	}

	return false
}

func (helper *DeWordContinuityHelper) GetLastWord(text string) string {
	words := regexp.MustCompile(`[\s\.]+`).Split(text, -1)
	return words[len(words)-1]
}

func NewDeutsch() *Deutsch {
	language := NewLanguage()
	language.Abbreviations = NewSetFromArray(append(DeutschAbbreviations, EnAbbreviations...))
	language.IsPunctuationBetweenQuotes = true
	language.WordContinuityHelper = &DeWordContinuityHelper{}
	return &Deutsch{
		Language: *language,
	}
}
