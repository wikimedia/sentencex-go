package languages_test

import (
	"strings"
	"testing"

	"github.com/wikimedia/sentencex-go/languages"
)

type SegmentationTest struct {
	text      string
	sentences []string
	skip      bool
}


func LanguageTest(t *testing.T, languageCode string, tests []SegmentationTest) {
	factory := languages.LanguageFactory{}
	language := factory.CreateLanguage(languageCode)
	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			if tt.skip {
				t.Skip()
			}
			segmented := language.Segment(tt.text)
			if len(segmented) != len(tt.sentences) {
				t.Errorf("Expected %d sentences, got %d", len(tt.sentences), len(segmented))
				t.Errorf("%#v", strings.Join(segmented, "❚"))
			} else {
				for i, actual_sentence := range segmented {
					if strings.TrimSpace(actual_sentence) != tt.sentences[i] && actual_sentence != tt.sentences[i]{
						t.Errorf("Expected '%s', got '%s'", tt.sentences[i], actual_sentence)
					}
				}
			}
		})
	}
}