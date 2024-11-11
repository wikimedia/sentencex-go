package languages

type Bulgarian struct {
	Language
}


var BulgarianAbbreviations = []string{
	"p.s",
	"акад",
	"ал",
	"б.р",
	"б.ред",
	"бел.а",
	"бел.пр",
	"бр",
	"бул",
	"в",
	"вж",
	"вкл",
	"вм",
	"вр",
	"г",
	"ген",
	"гр",
	"дж",
	"дм",
	"доц",
	"др",
	"ем",
	"заб",
	"зам",
	"инж",
	"к.с",
	"кв.м",
	"кв",
	"кг",
	"км",
	"кор",
	"куб.м",
	"куб",
	"л",
	"лв",
	"м.г",
	"м",
	"мин",
	"млн",
	"млрд",
	"мм",
	"н.с",
	"напр",
	"пл",
	"полк",
	"проф",
	"р",
	"рис",
	"с",
	"св",
	"сек",
	"см",
	"сп",
	"срв",
	"ст",
	"стр",
	"т.г",
	"т.е",
	"т.н",
	"т.нар",
	"т",
	"табл",
	"тел",
	"у",
	"ул",
	"фиг",
	"ха",
	"хил",
	"ч",
	"чл",
	"щ.д",
}

func NewBulgarian() *Bulgarian {
	language := NewLanguage()
	language.Abbreviations = NewSetFromArray(BulgarianAbbreviations)
	language.IsPunctuationBetweenQuotes = false

	return &Bulgarian{
		Language: *language,
	}
}