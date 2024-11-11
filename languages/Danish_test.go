package languages_test

import (
	"testing"
)

var datests = []SegmentationTest{
	{
		text:      "Hej Verden. Mit navn er Jonas.",
		sentences: []string{"Hej Verden.", "Mit navn er Jonas."},
	},
	{
		text:      "Hvad er dit navn? Mit nav er Jonas.",
		sentences: []string{"Hvad er dit navn?", "Mit nav er Jonas."},
	},
	{
		text:      "Lad os spørge Jane og co. De burde vide det.",
		sentences: []string{"Lad os spørge Jane og co.", "De burde vide det."},
		skip: 	true,
	},
	{
		text:      "De lukkede aftalen med Pitt, Briggs & Co. Det lukkede i går.",
		sentences: []string{"De lukkede aftalen med Pitt, Briggs & Co.", "Det lukkede i går."},
		skip: 	true,
	},
	{
		text:      "De holdt Skt. Hans i byen.",
		sentences: []string{"De holdt Skt. Hans i byen."},
	},
	{
		text:      "St. Michael's Kirke er på 5. gade nær ved lyset.",
		sentences: []string{"St. Michael's Kirke er på 5. gade nær ved lyset."},
	},
}

func TestDanish(t *testing.T) {
	LanguageTest(t, "da", datests)
}
