package languages_test

import (
	"testing"
)

var mltests = []struct {
	text      string
	sentences []string
	skip      bool
}{

	{
		text:      "Roses Are Red. Violets Are Blue",
		sentences: []string{"Roses Are Red.", "Violets Are Blue"},
	},
	{
		text:      "ഇത് ഡോ. ശിവൻ. ഇദ്ദേഹമാണ് ഞാൻ പറഞ്ഞയാൾ",
		sentences: []string{"ഇത് ഡോ. ശിവൻ.", "ഇദ്ദേഹമാണ് ഞാൻ പറഞ്ഞയാൾ"},
	},
	{
		text:      "ഇത് മി. കെ. പി. മോഹനൻ",
		sentences: []string{"ഇത് മി. കെ. പി. മോഹനൻ"},
	},
	{
		text:      "ഇത് പ്രൊ. കെ.പി. മോഹനൻ",
		sentences: []string{"ഇത് പ്രൊ. കെ.പി. മോഹനൻ"},
	},
	{
		text:      "ഇത് Dr. മോഹനൻ",
		sentences: []string{"ഇത് Dr. മോഹനൻ"},
	},
}

func TestMalayalam(t *testing.T) {
	LanguageTest(t, "ml", mltests)
}
