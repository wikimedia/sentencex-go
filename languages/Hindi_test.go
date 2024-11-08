package languages_test

import (
	"testing"
)

var hitests = []SegmentationTest{

	{
		text: "सच्चाई यह है कि इसे कोई नहीं जानता। हो सकता है यह फ़्रेन्को के खिलाफ़ कोई विद्रोह रहा हो, या फिर बेकाबू हो गया कोई आनंदोत्सव।",
		sentences: []string{
			"सच्चाई यह है कि इसे कोई नहीं जानता।",
			"हो सकता है यह फ़्रेन्को के खिलाफ़ कोई विद्रोह रहा हो, या फिर बेकाबू हो गया कोई आनंदोत्सव।",
		},
	},
}

func TestHindi(t *testing.T) {
	LanguageTest(t, "hi", hitests)
}
