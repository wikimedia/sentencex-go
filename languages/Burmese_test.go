package languages_test

import (
	"testing"
)

var mytests = []SegmentationTest{

	{
		text: "ခင္ဗ်ားနာမည္ဘယ္လိုေခၚလဲ။ င္ေနေကာင္းလား။",
		sentences: []string{
			"ခင္ဗ်ားနာမည္ဘယ္လိုေခၚလဲ။", "င္ေနေကာင္းလား။",
		},
	},
}

func TestMurmese(t *testing.T) {
	LanguageTest(t, "my", mytests)
}
