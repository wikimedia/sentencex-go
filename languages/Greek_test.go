package languages_test

import (
	"testing"
)

var eltests = []SegmentationTest{

	{
		text: "Με συγχωρείτε· πού είναι οι τουαλέτες; Τις Κυριακές δε δούλευε κανένας. το κόστος του σπιτιού ήταν £260.950,00.",
		sentences: []string{
			"Με συγχωρείτε· πού είναι οι τουαλέτες;",
			"Τις Κυριακές δε δούλευε κανένας.",
			"το κόστος του σπιτιού ήταν £260.950,00.",
		},
	},
}

func TestGreek(t *testing.T) {
	LanguageTest(t, "el", eltests)
}
