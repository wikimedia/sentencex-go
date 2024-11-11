package languages

type Gujarati struct {
	Language
}


var GujaratiAbbreviations = []string{
	"એ",
	"બી",
	"સી",
	"ડી",
	"ઈ",
	"એફ",
	"જી",
	"એચ",
	"આઈ",
	"જે",
	"કે",
	"એલ",
	"એમ",
	"એન",
	"ઓ",
	"પી",
	"ક્યૂ",
	"આર",
	"એસ",
	"ટી",
	"યૂ",
	"વી",
	"ડબલ્યૂ",
	"એક્સ",
	"વાય",
	"જેડ",
}

func NewGujarati() *Gujarati {
	language := NewLanguage()
	language.Abbreviations = NewSetFromArray(append(GujaratiAbbreviations, EnAbbreviations...))

	return &Gujarati{
		Language: *language,
	}
}
