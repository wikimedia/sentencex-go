package languages

type Arabic struct {
	Language
}

var ArabicAbbreviations = []string{
	"ا. د",
	"ا.د",
	"ا.ش.ا",
	"ا",
	"ت.ب",
	"ج.ب",
	"ج.م.ع",
	"جم",
	"س.ت",
	"سم",
	"ص.ب.",
	"ص.ب",
	"كج.",
	"كلم.",
	"م.ب",
	"م",
	"ه",
}

func NewArabic() *Arabic {
	language := NewLanguage()
	language.Abbreviations = NewSetFromArray(append(ArabicAbbreviations, EnAbbreviations...))
	return &Arabic{
		Language: *language,
	}
}
