package languages_test

import (
	"testing"
)

var urtests = []SegmentationTest{

	{
		text: "کیا حال ہے؟ ميرا نام ___ ەے۔ میں حالا تاوان دےدوں؟",
		sentences: []string{

			"کیا حال ہے؟", "ميرا نام ___ ەے۔", "میں حالا تاوان دےدوں؟",
		},
	},
}

func TestUrdu(t *testing.T) {
	LanguageTest(t, "ur", urtests)
}
