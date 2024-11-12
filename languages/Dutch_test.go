package languages_test

import (
	"testing"
)

var nltests = []SegmentationTest{
	{
		text:      "Hij schoot op de JP8-brandstof toen de Surface-to-Air (sam)-missiles op hem af kwamen. 81 procent van de schoten was raak.",
		sentences: []string{"Hij schoot op de JP8-brandstof toen de Surface-to-Air (sam)-missiles op hem af kwamen.", "81 procent van de schoten was raak."},
	},
	{
		text:      "81 procent van de schoten was raak. ...en toen barste de hel los.",
		sentences: []string{"81 procent van de schoten was raak.", "...", "en toen barste de hel los."},
	},
	{
		text:      "Afkorting aanw. vnw.",
		sentences: []string{"Afkorting aanw. vnw."},
	},
}

func TestDutch(t *testing.T) {
	LanguageTest(t, "nl", nltests)
}
