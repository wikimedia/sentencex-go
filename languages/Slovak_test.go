package languages_test

import (
	"testing"
)

var sktests = []SegmentationTest{
	{
		text:      "Ide o majiteľov firmy ABTrade s. r. o., ktorí stoja aj za ďalšími spoločnosťami, napr. XYZCorp a.s.",
		sentences: []string{"Ide o majiteľov firmy ABTrade s. r. o., ktorí stoja aj za ďalšími spoločnosťami, napr. XYZCorp a.s."},
	},
	{
		text:      "„Prieskumy beriem na ľahkú váhu. V podstate ma to nezaujíma,“ reagoval Matovič na prieskum agentúry Focus.",
		sentences: []string{"„Prieskumy beriem na ľahkú váhu. V podstate ma to nezaujíma,“ reagoval Matovič na prieskum agentúry Focus."},
	},
	{
		text:      "Toto sa mi podarilo až na 10. pokus, ale stálo to za to.",
		sentences: []string{"Toto sa mi podarilo až na 10. pokus, ale stálo to za to."},
	},
	{
		text:      "Ide o príslušníkov XII. Pluku špeciálneho určenia.",
		sentences: []string{"Ide o príslušníkov XII. Pluku špeciálneho určenia."},
	},
	{
		text:      "Spoločnosť bola založená 7. Apríla 2020, na zmluve však figuruje dátum 20. marec 2020.",
		sentences: []string{"Spoločnosť bola založená 7. Apríla 2020, na zmluve však figuruje dátum 20. marec 2020."},
	},
}

func TestSlovak(t *testing.T) {
	LanguageTest(t, "sk", sktests)
}
