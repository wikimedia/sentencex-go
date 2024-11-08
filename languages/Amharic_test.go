package languages_test

import (
	"testing"
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
	LanguageTest(t, "am", amtests)
}
