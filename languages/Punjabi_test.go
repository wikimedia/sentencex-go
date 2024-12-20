package languages_test

import (
	"testing"
)

var patests = []SegmentationTest{

	{
		text: "ਸਰੋਜਿਨੀ ਨਾਇਡੂ ਦਾ ਜਨਮ 13 ਫਰਵਰੀ 1879 ਨੂੰ ਭਾਰਤ ਦੇ ਸ਼ਹਿਰ ਹੈਦਰਾਬਾਦ ਵਿੱਚ ਹੋਇਆ ਸੀ। ਉਸ ਦੇ ਪਿਤਾ ਅਘੋਰਨਾਥ ਚੱਟੋਪਾਧਿਆਏ ਇੱਕ ਨਾਮੀ ਵਿਦਵਾਨ ਅਤੇ ਮਾਂ ਬਰਾਦਾ ਸੁੰਦਰੀ ਦੇਬੀ ਕਵਿਤਰੀ ਸੀ ਅਤੇ ਬੰਗਾਲੀ ਵਿੱਚ ਲਿਖਦੀ ਸੀ।",
		sentences: []string{"ਸਰੋਜਿਨੀ ਨਾਇਡੂ ਦਾ ਜਨਮ 13 ਫਰਵਰੀ 1879 ਨੂੰ ਭਾਰਤ ਦੇ ਸ਼ਹਿਰ ਹੈਦਰਾਬਾਦ ਵਿੱਚ ਹੋਇਆ ਸੀ।",
			"ਉਸ ਦੇ ਪਿਤਾ ਅਘੋਰਨਾਥ ਚੱਟੋਪਾਧਿਆਏ ਇੱਕ ਨਾਮੀ ਵਿਦਵਾਨ ਅਤੇ ਮਾਂ ਬਰਾਦਾ ਸੁੰਦਰੀ ਦੇਬੀ ਕਵਿਤਰੀ ਸੀ ਅਤੇ ਬੰਗਾਲੀ ਵਿੱਚ ਲਿਖਦੀ ਸੀ।"},
	},
}

func TestPunjabi(t *testing.T) {
	LanguageTest(t, "pa", patests)
}
