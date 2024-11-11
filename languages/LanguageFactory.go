package languages

type LanguageFactory struct{}

func  (f *LanguageFactory) CreateLanguage(language string) ILanguage {

	switch language {
	case "am":
		return NewAmharic()
	case "ar":
		return NewAmharic()
	case "bg":
		return NewBulgarian()
	case "da":
		return NewDanish()
	case "de":
		return NewDeutsch()
	case "el":
		return NewGreek()
	case "en":
		return NewEnglish()
	case "es":
		return NewSpanish()
	case "fi":
		return NewFinnish()
	case "fr":
		return NewFrench()
	case "gu":
		return NewGujarati()
	case "hi":
		return NewHindi()
	case "hy":
		return NewArmenian()
	case "it":
		return NewItalian()
	case "ml":
		return NewMalayalam()
	case "ru":
		return NewRussian()
	case "ta":
		return NewTamil()
}

	if LANGUAGE_FALLBACKS[language] != nil {
		for _, fallback := range LANGUAGE_FALLBACKS[language] {
			return f.CreateLanguage(fallback)
		}
	}
	return NewLanguage()
}
