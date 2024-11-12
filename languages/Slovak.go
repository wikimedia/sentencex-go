package languages

import (
	"regexp"
	"strings"
)

type Slovak struct {
	Language
}

var SK_MONTHS = []string{
	"Január",
	"Február",
	"Marec",
	"Apríl",
	"Máj",
	"Jún",
	"Júl",
	"August",
	"September",
	"Október",
	"November",
	"December",
	"Januára",
	"Februára",
	"Marca",
	"Apríla",
	"Mája",
	"Júna",
	"Júla",
	"Augusta",
	"Septembra",
	"Októbra",
	"Novembra",
	"Decembra",
}
var SlovakAbbreviations = []string{
	"a. d",
	"a. g. p",
	"a. i. i",
	"a. k. a",
	"a. m",
	"a. r. k",
	"a. s. a. p",
	"a. s",
	"a. v",
	"a.d",
	"a.g.p",
	"a.i.i",
	"a.k.a",
	"a.m",
	"a.s.a.p",
	"a.s",
	"a.v",
	"akad",
	"al",
	"apod",
	"arm",
	"atď.",
	"atd",
	"atď",
	"bc",
	"bros",
	"c. k",
	"c.k",
	"č",
	"cca",
	"co",
	"corp",
	"čs",
	"csc",
	"čsl",
	"d. c",
	"d.c",
	"doc",
	"dr",
	"drsc",
	"e. t",
	"e.t",
	"el",
	"etc",
	"ev",
	"gen",
	"hl",
	"hod",
	"i. b",
	"i.b",
	"ii",
	"iii",
	"inc",
	"ind",
	"ing",
	"iv",
	"jr",
	"judr",
	"k. o",
	"k.o",
	"kol",
	"konkr",
	"kt",
	"ll. m",
	"ll.m",
	"ltd",
	"m. n. m",
	"m.n.m",
	"m.o",
	"max",
	"mgr",
	"mil",
	"min",
	"ml",
	"mld",
	"mr",
	"mudr",
	"mvdr",
	"n. a",
	"n. o",
	"n. w. a",
	"n.a",
	"n.o",
	"n.w.a",
	"nám",
	"napr",
	"např",
	"naprk",
	"nár",
	"nešp",
	"no",
	"nr",
	"o. c. p",
	"o. f. i",
	"o. k",
	"o. z",
	"o.c.p",
	"o.f.i",
	"o.i",
	"o.k",
	"o.z",
	"obr",
	"obv",
	"odd",
	"ods",
	"os",
	"p. a",
	"p. n. l",
	"p. s",
	"p.a",
	"p.n.l",
	"p.s",
	"p",
	"paeddr",
	"pedg",
	"ph. d",
	"ph.d",
	"phd",
	"phdr",
	"písm",
	"plgr",
	"pod",
	"pok",
	"pol. pr",
	"pol.pr",
	"por",
	"pozn",
	"pp",
	"pr",
	"prek",
	"príp",
	"prof",
	"r. o",
	"r.o",
	"red",
	"resp",
	"rndr",
	"roz",
	"rozh",
	"rsdr",
	"rtg",
	"s. a",
	"s. e. g",
	"š. p",
	"s. r. o",
	"s.a",
	"s.e.g",
	"š.p",
	"s.r.o",
	"skr",
	"sl",
	"slov",
	"soc",
	"sp",
	"spol",
	"sr",
	"st",
	"št",
	"stor",
	"str",
	"stred",
	"súkr",
	"sv",
	"sz",
	"t. č",
	"t. j",
	"t. z",
	"t.č",
	"t.j",
	"t.z",
	"tel",
	"tis",
	"tj",
	"tr",
	"tu",
	"tvz",
	"tz",
	"tzn",
	"tzv",
	"ú. p. v. o",
	"u. s",
	"ú.p.v.o",
	"u.s",
	"ul",
	"v. sp",
	"v.sp",
	"var",
	"vi",
	"viď",
	"vs",
	"vyd",
	"vz",
	"xx",
	"z. z",
	"z.z",
	"zák",
	"zb",
	"zdravot",
	"zs",
	"zz",
}

type SkWordContinuityHelper struct {
}

func (helper *SkWordContinuityHelper) ContinueInNextWord(textAfterBoundary string) bool {
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

	for _, month := range SK_MONTHS {
		if nextWord == month || (strings.ToUpper(string(nextWord[0]))+nextWord[1:] == month) {
			return true
		}
	}

	return false

}

func (helper *SkWordContinuityHelper) GetLastWord(text string) string {
	words := regexp.MustCompile(`[\s\.]+`).Split(text, -1)
	return words[len(words)-1]
}

func NewSlovak() *Slovak {
	language := NewLanguage()
	language.Abbreviations = NewSetFromArray(append(SlovakAbbreviations, strings.Fields(roman_numerals+" "+strings.ToUpper(roman_numerals))...))
	language.WordContinuityHelper = &SkWordContinuityHelper{}
	return &Slovak{
		Language: *language,
	}
}
