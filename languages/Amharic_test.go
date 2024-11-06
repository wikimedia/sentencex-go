package languages_test

import (
	"strings"
	"testing"

	"github.com/wikimedia/sentencex-go/languages"
)

var amtests = []struct {
	text      string
	sentences []string
	skip      bool
}{

	{
		text:       "እንደምን አለህ፧መልካም ቀን ይሁንልህ።እባክሽ ያልሽዉን ድገሚልኝ።",
		sentences: []string{"እንደምን አለህ፧", "መልካም ቀን ይሁንልህ።", "እባክሽ ያልሽዉን ድገሚልኝ።"},
	},
	{
		text:      "ቴዎድሮስ ጥር ፮ ቀን ፲፰፻፲፩ ዓ.ም. ሻርጌ በተባለ ቦታ ቋራ ውስጥ፣ ከጎንደር ከተማ በስተ ምዕራብ ተወለዱ።",
		sentences: []string{"ቴዎድሮስ ጥር ፮ ቀን ፲፰፻፲፩ ዓ.ም. ሻርጌ በተባለ ቦታ ቋራ ውስጥ፣ ከጎንደር ከተማ በስተ ምዕራብ ተወለዱ።"},
	},
}

func TestAmharic(t *testing.T) {
	factory := languages.LanguageFactory{}
	language := factory.CreateLanguage("am")
	for _, tt := range amtests {
		t.Run(tt.text, func(t *testing.T) {
			if tt.skip {
				t.Skip()
			}
			segmented := language.Segment(tt.text)
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
