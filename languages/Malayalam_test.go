package languages_test

import (
	"strings"
	"testing"

	"github.com/wikimedia/sentencex-go/languages"
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
	factory := languages.LanguageFactory{}
	english := factory.CreateLanguage("ml")
	for _, tt := range mltests {
		t.Run(tt.text, func(t *testing.T) {
			if tt.skip {
				t.Skip()
			}
			segmented := english.Segment(tt.text)
			if len(segmented) != len(tt.sentences) {
				t.Errorf("Expected %d sentences, got %d", len(tt.sentences), len(segmented))
				t.Error(segmented)
			} else {
				for i, actual_sentence := range segmented {
					if strings.TrimSpace(actual_sentence) != tt.sentences[i] {
						t.Errorf("Expected '%s', got '%s'", tt.sentences[i], actual_sentence)
					}
				}
			}
		})
	}
}
