package languages

type Amharic struct {
	Language
}

var AmharicAbbreviations = []string{
	"ዓ",
	"ም",
}


func NewAmharic() *Amharic {
	language := NewLanguage()
	language.Abbreviations = NewSetFromArray(append(AmharicAbbreviations, EnAbbreviations...))

	return &Amharic{
		Language: *language,
	}
}
