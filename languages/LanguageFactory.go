package languages

type LanguageFactory struct{}

func  (f *LanguageFactory) CreateLanguage(language string) ILanguage {

	switch language {
	case "en":
		return NewEnglish()
	case "ml":
		return NewMalayalam()
	}

	if LANGUAGE_FALLBACKS[language] != nil {
		for _, fallback := range LANGUAGE_FALLBACKS[language] {
			return f.CreateLanguage(fallback)
		}
	}
	return NewLanguage()
}
