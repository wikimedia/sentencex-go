package languages

type Telugu struct {
	Language
}

var TeluguAbbreviations = []string{
	"ఎ",
	"బి",
	"సి",
	"డి",
	"ఈ",
	"ఎఫ్",
	"జి",
	"హెచ్",
	"ఐ",
	"జె",
	"కె",
	"ఎల్",
	"ఎం",
	"ఎన్",
	"ఓ",
	"పి",
	"క్యూ",
	"ఆర్",
	"ఎస్",
	"టి",
	"యూ",
	"వి",
	"డబ్ల్యూ",
	"ఎక్స్",
	"వై",
	"జెడ్",
}

func NewTelugu() *Telugu {
	language := NewLanguage()
	language.Abbreviations = NewSetFromArray(TeluguAbbreviations)
	return &Telugu{
		Language: *language,
	}
}
