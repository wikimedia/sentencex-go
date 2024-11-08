package languages

type Tamil struct {
	Language
}

var vowelSigns = []string{"ா", "ி", "ீ", "ு", "ூ", "ெ", "ே", "ை", "ொ", "ோ", "ௌ"}
var vowels = []string{"அ", "ஆ", "இ", "ஈ", "உ", "ஊ", "எ", "ஏ", "ஐ", "ஒ", "ஓ", "ஔ"}
var consonants = []string{
	"க", "ங", "ச", "ஞ", "ட", "ண", "த", "ந", "ப", "ம", "ய", "ர", "ல", "வ", "ழ", "ள", "ற", "ன",
}

var consonantVowels = []string{}

func init() {
	for _, consonant := range consonants {
		for _, vowelSign := range vowelSigns {
			consonantVowels = append(consonantVowels, consonant+vowelSign)
		}
	}
}

var TamilAbbreviations = []string{
	"ஏ",
	"பி",
	"சி",
	"டி",
	"ஈ",
	"எஃப்",
	"ஜி",
	"ஹேச்",
	"ஐ",
	"ஜே",
	"கே",
	"எல்",
	"எம்",
	"என்",
	"ஓ",
	// "பி",
	"க்யூ",
	"ஆர்",
	"எஸ்",
	// "டி",
	"யூ",
	"வி",
	"டபிள்யூ",
	"எக்ஸ்",
	"வை",
	"ஜெட்",

}



func NewTamil() *Tamil {
	language := NewLanguage()
	TamilAbbreviations = append(TamilAbbreviations, vowels...)
	TamilAbbreviations = append(TamilAbbreviations, consonants...)
	TamilAbbreviations = append(TamilAbbreviations, consonantVowels...)
	language.Abbreviations = NewSetFromArray(append(TamilAbbreviations, EnAbbreviations...  ))
	language.IsPunctuationBetweenQuotes = true

	return &Tamil{
		Language: *language,
	}
}
