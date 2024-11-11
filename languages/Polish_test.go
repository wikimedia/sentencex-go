package languages_test

import (
	"testing"
)

var pltests = []SegmentationTest{
	{
		text:      "To słowo bałt. jestskrótem.",
		sentences: []string{"To słowo bałt. jestskrótem."},
	},
}

func TestPolish(t *testing.T) {
	LanguageTest(t, "pl", pltests)
}
