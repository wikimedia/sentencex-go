package languages

type Hindi struct {
	Language
}


var HindiAbbreviations = []string{
	"ए",
	"बी",
	"सी",
	"डी",
	"ई",
	"एफ",
	"जी",
	"एच",
	"आई",
	"जे",
	"के",
	"एल",
	"एम",
	"एन",
	"ओ",
	"पी",
	"क्यू",
	"आर",
	"एस",
	"टी",
	"यू",
	"भी",
	"डब्लू",
	"एक्स",
	"वाई",
	"जेड",
}

func NewHindi() *Hindi {
	language := NewLanguage()
	language.Abbreviations = NewSetFromArray(append(HindiAbbreviations, EnAbbreviations...))
	language.IsPunctuationBetweenQuotes = true

	return &Hindi{
		Language: *language,
	}
}
