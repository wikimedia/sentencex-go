package languages_test

import (
	"testing"
)

var kktests = []SegmentationTest{
	{
		text: "Мұхитқа тікелей шыға алмайтын мемлекеттердің ішінде Қазақстан - ең үлкені.",
		sentences: []string{
			"Мұхитқа тікелей шыға алмайтын мемлекеттердің ішінде Қазақстан - ең үлкені.",
		},
	},
	{
		text: "Оқушылар үйі, Достық даңғылы, Абай даналығы, ауыл шаруашылығы – кім? не?",
		sentences: []string{
			"Оқушылар үйі, Достық даңғылы, Абай даналығы, ауыл шаруашылығы – кім?",
			"не?",
		},
		skip: true,
	},
	{
		text: "Әр түрлі өлшемнің атауы болып табылатын м (метр), см (сантиметр), кг (киллограмм), т (тонна), га (гектар), ц (центнер), т. б. (тағы басқа), тәрізді белгілер де қысқарған сөздер болып табылады.",
		sentences: []string{
			"Әр түрлі өлшемнің атауы болып табылатын м (метр), см (сантиметр), кг (киллограмм), т (тонна), га (гектар), ц (центнер), т. б. (тағы басқа), тәрізді белгілер де қысқарған сөздер болып табылады.",
		},
	},
	{
		text: "Мысалы: обкомға (облыстық комитетке) барды, ауаткомда (аудандық атқару комитетінде) болды, педучилищеге (педагогтік училищеге) түсті, медпункттің (медициналық пункттің) алдында т. б.",
		sentences: []string{
			"Мысалы: обкомға (облыстық комитетке) барды, ауаткомда (аудандық атқару комитетінде) болды, педучилищеге (педагогтік училищеге) түсті, медпункттің (медициналық пункттің) алдында т. б.",
		},
	},
	{
		text: "Елдің жалпы ішкі өнімі ЖІӨ (номинал) = $225.619 млрд (2014)",
		sentences: []string{
			"Елдің жалпы ішкі өнімі ЖІӨ (номинал) = $225.619 млрд (2014)",
		},
	},
	{
		text: "Ресейдiң әлеуметтiк-экономикалық жағдайы.XVIII ғасырдың бiрiншi ширегiнде Ресейге тән нәрсе.",
		sentences: []string{
			"Ресейдiң әлеуметтiк-экономикалық жағдайы.",
			"XVIII ғасырдың бiрiншi ширегiнде Ресейге тән нәрсе.",
		},
	},
	{
		text: "(«Егемен Қазақстан», 7 қыркүйек 2012 жыл. №590-591); Бұл туралы кеше санпедқадағалау комитетінің облыыстық департаменті хабарлады. («Айқын», 23 сəуір 2010 жыл. № 70).",
		sentences: []string{
			"(«Егемен Қазақстан», 7 қыркүйек 2012 жыл. №590-591); Бұл туралы кеше санпедқадағалау комитетінің облыыстық департаменті хабарлады. («Айқын», 23 сəуір 2010 жыл. № 70).",
		},
	},
	{
		text: "Иран революциясы (1905 — 11) және азаматтық қозғалыс (1918 — 21) кезінде А. Фарахани, М. Кермани, М. Т. Бехар, т.б. ақындар демократиялық идеяның жыршысы болды.",
		sentences: []string{
			"Иран революциясы (1905 — 11) және азаматтық қозғалыс (1918 — 21) кезінде А. Фарахани, М. Кермани, М. Т. Бехар, т.б. ақындар демократиялық идеяның жыршысы болды.",
		},
	},
	{
		text: "Владимир Федосеев: Аттар магиясы енді жоқ http://www.vremya.ru/2003/179/10/80980.html",
		sentences: []string{
			"Владимир Федосеев: Аттар магиясы енді жоқ http://www.vremya.ru/2003/179/10/80980.html",
		},
	},
	{
		text: "Бірақ оның енді не керегі бар? — деді.",
		sentences: []string{
			"Бірақ оның енді не керегі бар? — деді.",
		},
	},
	{
		text: "Сондықтан шапаныма жегізіп отырғаным! - деп, жауап береді.",
		sentences: []string{
			"Сондықтан шапаныма жегізіп отырғаным! - деп, жауап береді.",
		},
	},
	{
		text: "Б.з.б. 6 – 3 ғасырларда конфуцийшілдік, моизм, легизм мектептерінің қалыптасуы нәтижесінде Қытай философиясы пайда болды.",
		sentences: []string{
			"Б.з.б. 6 – 3 ғасырларда конфуцийшілдік, моизм, легизм мектептерінің қалыптасуы нәтижесінде Қытай философиясы пайда болды.",
		},
	},
	{
		text: "'Та марбута' тек сөз соңында екі түрде жазылады:",
		sentences: []string{
			"'Та марбута' тек сөз соңында екі түрде жазылады:",
		},
	},
}

func TestKazhakh(t *testing.T) {
	LanguageTest(t, "kk", kktests)
}
