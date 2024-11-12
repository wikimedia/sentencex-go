package languages_test

import (
	"testing"
)

var zhtests = []SegmentationTest{

	{
		text: "کیا حال ہے؟ ميرا نام ___ ەے۔ میں حالا تاوان دےدوں؟",
		sentences: []string{

			"کیا حال ہے؟", "ميرا نام ___ ەے۔", "میں حالا تاوان دےدوں؟",
		},
	},
}

func TestChinese(t *testing.T) {
	LanguageTest(t, "zh", zhtests)
}
