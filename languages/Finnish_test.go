package languages_test

import (
	"strings"
	"testing"

	"github.com/wikimedia/sentencex-go/languages"
)

var fitests = []struct {
	text      string
	sentences []string
	skip      bool
}{
	{
		text:      "Se julkaistiin singlenä 7. heinäkuuta 1997, ja se nousi listaykköseksi yhtyeen kotimaassa Britanniassa sekä Irlannissa, Suomessa, Espanjassa ja Kanadassa",
		sentences: []string{"Se julkaistiin singlenä 7. heinäkuuta 1997, ja se nousi listaykköseksi yhtyeen kotimaassa Britanniassa sekä Irlannissa, Suomessa, Espanjassa ja Kanadassa"},
	},
	{
		text:      "Brittiläinen musiikkilehti NME valitsi lokakuussa 2011 ”D’You Know What I Meanin?” sijalle 77 listallaan, joka sisälsi 150 parasta kappaletta vuosilta 1996–2011.",
		sentences: []string{"Brittiläinen musiikkilehti NME valitsi lokakuussa 2011 ”D’You Know What I Meanin?” sijalle 77 listallaan, joka sisälsi 150 parasta kappaletta vuosilta 1996–2011."},
	},
	{
		text:      "Netistä ladattu musiikki on otettu huomioon singlelistalla 3. lokakuuta 2007 lähtien.[13] Uudistus muutti listan luonnetta.",
		sentences: []string{"Netistä ladattu musiikki on otettu huomioon singlelistalla 3. lokakuuta 2007 lähtien.[13]", "Uudistus muutti listan luonnetta."},
	},
	{
		text:      "Radiomafia oli Yleisradion pääasiassa nuorille ja nuorille aikuisille suunnattu radiokanava, joka aloitti toimintansa vuoden 1990 radiouudistuksessa 1. kesäkuuta 1990 ja lopetti Ylen radiouudistuksen myötä 12. tammikuuta 2003.",
		sentences: []string{"Radiomafia oli Yleisradion pääasiassa nuorille ja nuorille aikuisille suunnattu radiokanava, joka aloitti toimintansa vuoden 1990 radiouudistuksessa 1. kesäkuuta 1990 ja lopetti Ylen radiouudistuksen myötä 12. tammikuuta 2003."},
	},
	{
		text:      "Dr. Alban (oikealta nimeltä Alban Uzoma Nwapa, s. 26. elokuuta 1957 Oguta, Brittiläinen Nigeria) on nigerialaissyntyinen ruotsalainen eurodance/rap/reggae -artisti.",
		sentences: []string{"Dr. Alban (oikealta nimeltä Alban Uzoma Nwapa, s. 26. elokuuta 1957 Oguta, Brittiläinen Nigeria) on nigerialaissyntyinen ruotsalainen eurodance/rap/reggae -artisti."},
	},
	{
		text:      "Hän syntyi 12. joulukuuta 1980 Helsingissä. Hän on suomalainen näyttelijä.",
		sentences: []string{"Hän syntyi 12. joulukuuta 1980 Helsingissä.", "Hän on suomalainen näyttelijä."},
	},
	{
		text:      "Vuonna 2005 hän voitti ensimmäisen palkintonsa. Se oli merkittävä hetki hänen urallaan.",
		sentences: []string{"Vuonna 2005 hän voitti ensimmäisen palkintonsa.", "Se oli merkittävä hetki hänen urallaan."},
	},
	{
		text:      "Hän asuu Helsingissä, mutta työskentelee usein ulkomailla. Hänen työnsä vie hänet usein eri puolille maailmaa.",
		sentences: []string{"Hän asuu Helsingissä, mutta työskentelee usein ulkomailla.", "Hänen työnsä vie hänet usein eri puolille maailmaa."},
	},
	{
		text:      "Hän on kirjoittanut useita kirjoja. Hänen viimeisin teoksensa julkaistiin vuonna 2020.",
		sentences: []string{"Hän on kirjoittanut useita kirjoja.", "Hänen viimeisin teoksensa julkaistiin vuonna 2020."},
	},
}

func TestFinnish(t *testing.T) {
	factory := languages.LanguageFactory{}
	language := factory.CreateLanguage("de")
	for _, tt := range fitests {
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
