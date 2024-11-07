package languages_test

import (
	"strings"
	"testing"

	"github.com/wikimedia/sentencex-go/languages"
)

var artests = []struct {
	text      string
	sentences []string
	skip      bool
}{

	{
		text:      "እንደምን አለህ፧መልካም ቀን ይሁንልህ።እባክሽ ያልሽዉን ድገሚልኝ።",
		sentences: []string{"እንደምን አለህ፧", "መልካም ቀን ይሁንልህ።", "እባክሽ ያልሽዉን ድገሚልኝ።"},
	},
	{
		text:      "ቴዎድሮስ ጥር ፮ ቀን ፲፰፻፲፩ ዓ.ም. ሻርጌ በተባለ ቦታ ቋራ ውስጥ፣ ከጎንደር ከተማ በስተ ምዕራብ ተወለዱ።",
		sentences: []string{"ቴዎድሮስ ጥር ፮ ቀን ፲፰፻፲፩ ዓ.ም. ሻርጌ በተባለ ቦታ ቋራ ውስጥ፣ ከጎንደር ከተማ በስተ ምዕራብ ተወለዱ።"},
	},
	{
		text: "سؤال وجواب: ماذا حدث بعد الانتخابات الايرانية؟ طرح الكثير من التساؤلات غداة ظهور نتائج الانتخابات الرئاسية الايرانية التي أججت مظاهرات واسعة واعمال عنف بين المحتجين على النتائج ورجال الامن. يقول معارضو الرئيس الإيراني إن الطريقة التي اعلنت بها النتائج كانت مثيرة للاستغراب.",
		sentences: []string{
			"سؤال وجواب:",
			"ماذا حدث بعد الانتخابات الايرانية؟",
			"طرح الكثير من التساؤلات غداة ظهور نتائج الانتخابات الرئاسية الايرانية التي أججت مظاهرات واسعة واعمال عنف بين المحتجين على النتائج ورجال الامن.",
			"يقول معارضو الرئيس الإيراني إن الطريقة التي اعلنت بها النتائج كانت مثيرة للاستغراب.",
		},
		skip: true,
	},
	// has unicode character. Need arabic language knowledge to resolve the issue
	{
		text: "وقال د‪.‬ ديفيد ريدي و الأطباء الذين كانوا يعالجونها في مستشفى برمنجهام إنها كانت تعاني من أمراض أخرى. وليس معروفا ما اذا كانت قد توفيت بسبب اصابتها بأنفلونزا الخنازير.",
		sentences: []string{
			"وقال د‪.‬ ديفيد ريدي و الأطباء الذين كانوا يعالجونها في مستشفى برمنجهام إنها كانت تعاني من أمراض أخرى.",
			"وليس معروفا ما اذا كانت قد توفيت بسبب اصابتها بأنفلونزا الخنازير.",
		},
		skip: true,
	},
	{
		text: "ومن المنتظر أن يكتمل مشروع خط أنابيب نابوكو البالغ طوله 3300 كليومترا في 12‪/‬08‪/‬2014 بتكلفة تُقدر بـ 7.9 مليارات يورو أي نحو 10.9 مليارات دولار. ومن المقرر أن تصل طاقة ضخ الغاز في المشروع 31 مليار متر مكعب انطلاقا من بحر قزوين مرورا بالنمسا وتركيا ودول البلقان دون المرور على الأراضي الروسية.",
		sentences: []string{
			"ومن المنتظر أن يكتمل مشروع خط أنابيب نابوكو البالغ طوله 3300 كليومترا في 12‪/‬08‪/‬2014 بتكلفة تُقدر بـ 7.9 مليارات يورو أي نحو 10.9 مليارات دولار.",
			"ومن المقرر أن تصل طاقة ضخ الغاز في المشروع 31 مليار متر مكعب انطلاقا من بحر قزوين مرورا بالنمسا وتركيا ودول البلقان دون المرور على الأراضي الروسية.",
		},
	},
	{
		text: "الاحد, 21 فبراير/ شباط, 2010, 05:01 GMT الصنداي تايمز: رئيس الموساد قد يصبح ضحية الحرب السرية التي شتنها بنفسه. العقل المنظم هو مئير داجان رئيس الموساد الإسرائيلي الذي يشتبه بقيامه باغتيال القائد الفلسطيني في حركة حماس محمود المبحوح في دبي.",
		sentences: []string{
			"الاحد, 21 فبراير/ شباط, 2010, 05:01 GMT الصنداي تايمز:",
			"رئيس الموساد قد يصبح ضحية الحرب السرية التي شتنها بنفسه.",
			"العقل المنظم هو مئير داجان رئيس الموساد الإسرائيلي الذي يشتبه بقيامه باغتيال القائد الفلسطيني في حركة حماس محمود المبحوح في دبي.",
		},
		skip: true,
	},
	{
		text: "عثر في الغرفة على بعض أدوية علاج ارتفاع ضغط الدم، والقلب، زرعها عملاء الموساد كما تقول مصادر إسرائيلية، وقرر الطبيب أن الفلسطيني قد توفي وفاة طبيعية ربما إثر نوبة قلبية، وبدأت مراسم الحداد عليه",
		sentences: []string{
			"عثر في الغرفة على بعض أدوية علاج ارتفاع ضغط الدم، والقلب،",
			"زرعها عملاء الموساد كما تقول مصادر إسرائيلية،",
			"وقرر الطبيب أن الفلسطيني قد توفي وفاة طبيعية ربما إثر نوبة قلبية،",
			"وبدأت مراسم الحداد عليه",
		},
		skip: true,
	},
}

func TestArabic(t *testing.T) {
	factory := languages.LanguageFactory{}
	language := factory.CreateLanguage("am")
	for _, tt := range artests {
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