package languages

type LanguageFactory struct{}

func  (f *LanguageFactory) CreateLanguage(language string) ILanguage {

	switch language {
	case "de":
		return NewDeutsch()
	case "en":
		return NewEnglish()
	case "am":
		return NewAmharic()
	case "ar":
		return NewAmharic()
	case "ml":
		return NewMalayalam()
	case "es":
		return NewSpanish()
	case "fi":
		return NewFinnish()
}

	if LANGUAGE_FALLBACKS[language] != nil {
		for _, fallback := range LANGUAGE_FALLBACKS[language] {
			return f.CreateLanguage(fallback)
		}
	}
	return NewLanguage()
}
