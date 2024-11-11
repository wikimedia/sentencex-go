package languages

type Kannada struct {
	Language
}

var KannadaAbbreviations = []string{
	"ಎ",
	"ಬಿ",
	"ಸಿ",
	"ಡಿ",
	"ಈ",
	"ಎಫ್",
	"ಜಿ",
	"ಹೆಚ್",
	"ಐ",
	"ಜೆ",
	"ಕೆ",
	"ಎಲ್",
	"ಎಂ",
	"ಎನ್",
	"ಓ",
	"ಪಿ",
	"ಕ್ಯೂ",
	"ಆರ್",
	"ಎಸ್",
	"ಟಿ",
	"ಯೂ",
	"ವಿ",
	"ಡಬಲ್ಯೂ",
	"ಎಕ್ಸ್",
	"ವೈ",
	"ಜೆಡ್",
}

func NewKannada() *Kannada {
	language := NewLanguage()
	language.Abbreviations = NewSetFromArray(append(KannadaAbbreviations, EnAbbreviations...))

	return &Kannada{
		Language: *language,
	}
}
