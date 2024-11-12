package languages

type Punjabi struct {
	Language
}

var PunjabiAbbreviations = []string{
	"ਏ",
	"ਬੀ",
	"ਸੀ",
	"ਡੀ",
	"ਈ",
	"ਐਫ",
	"ਜੀ",
	"ਐਚ",
	"ਆਈ",
	"ਜੇ",
	"ਕੇ",
	"ਐਲ",
	"ਐਮ",
	"ਐਨ",
	"ਓ",
	"ਪੀ",
	"ਕਿਊ",
	"ਆਰ",
	"ਐਸ",
	"ਟੀ",
	"ਯੂ",
	"ਵੀ",
	"ਡਬਲਯੂ",
	"ਐਕਸ",
	"ਵਾਈ",
	"ਜੇਡ",
}

func NewPunjabi() *Punjabi {
	language := NewLanguage()
	language.Abbreviations = NewSetFromArray(append(PunjabiAbbreviations, EnAbbreviations...))

	return &Punjabi{
		Language: *language,
	}
}
