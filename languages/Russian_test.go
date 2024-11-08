package languages_test

import (
	"strings"
	"testing"

	"github.com/wikimedia/sentencex-go/languages"
)

var rutests = []struct {
	text      string
	sentences []string
	skip      bool
}{
	{
		text: "Объем составляет 5 куб.м.",
		sentences: []string{
			"Объем составляет 5 куб.м.",
		},
	},
	{
		text: "Маленькая девочка бежала и кричала: «Не видали маму?».",
		sentences: []string{
			"Маленькая девочка бежала и кричала: «Не видали маму?».",
		},
	},
	{
		text: "Сегодня 27.10.14",
		sentences: []string{
			"Сегодня 27.10.14",
		},
	},
	{
		text: "«Я приду поздно»,  — сказал Андрей.",
		sentences: []string{
			"«Я приду поздно»,  — сказал Андрей.",
		},
	},
	{
		text: "«К чему ты готовишься? – спросила мама. – Завтра ведь выходной».",
		sentences: []string{
			"«К чему ты готовишься? – спросила мама. – Завтра ведь выходной».",
		},
	},
	{
		text: "По словам Пушкина, «Привычка свыше дана, замена счастью она».",
		sentences: []string{
			"По словам Пушкина, «Привычка свыше дана, замена счастью она».",
		},
	},
	{
		text: "Он сказал: «Я очень устал», и сразу же замолчал.",
		sentences: []string{
			"Он сказал: «Я очень устал», и сразу же замолчал.",
		},
	},
	{
		text: "Мне стало как-то ужасно грустно в это мгновение; однако что-то похожее на смех зашевелилось в душе моей.",
		sentences: []string{
			"Мне стало как-то ужасно грустно в это мгновение; однако что-то похожее на смех зашевелилось в душе моей.",
		},
	},
	{
		text: "Шухов как был в ватных брюках, не снятых на ночь (повыше левого колена их тоже был пришит затасканный, погрязневший лоскут, и на нем выведен черной, уже поблекшей краской номер Щ-854), надел телогрейку…",
		sentences: []string{
			"Шухов как был в ватных брюках, не снятых на ночь (повыше левого колена их тоже был пришит затасканный, погрязневший лоскут, и на нем выведен черной, уже поблекшей краской номер Щ-854), надел телогрейку…",
		},
	},
	{
		text: "Слово «дом» является синонимом жилища",
		sentences: []string{
			"Слово «дом» является синонимом жилища",
		},
	},
	{
		text: "В Санкт-Петербург на гастроли приехал театр «Современник»",
		sentences: []string{
			"В Санкт-Петербург на гастроли приехал театр «Современник»",
		},
	},
	{
		text: "Машина едет со скоростью 100 км/ч.",
		sentences: []string{
			"Машина едет со скоростью 100 км/ч.",
		},
	},
	{
		text: "Я поем и/или лягу спать.",
		sentences: []string{
			"Я поем и/или лягу спать.",
		},
	},
	{
		text: "Он не мог справиться с примером \"3 + (14:7) = 5\"",
		sentences: []string{
			"Он не мог справиться с примером \"3 + (14:7) = 5\"",
		},
	},
	{
		text: "Вот список: 1.мороженое, 2.мясо, 3.рис.",
		sentences: []string{
			"Вот список: 1.мороженое, 2.мясо, 3.рис.",
		},
	},
	{
		text: "Квартира 234 находится на 4-ом этаже.",
		sentences: []string{
			"Квартира 234 находится на 4-ом этаже.",
		},
	},
	{
		text: "В это время года температура может подниматься до 40°C.",
		sentences: []string{
			"В это время года температура может подниматься до 40°C.",
		},
	},
	{
		text: "Объем составляет 5м³.",
		sentences: []string{
			"Объем составляет 5м³.",
		},
	},
	{
		text: "Площадь комнаты 14м².",
		sentences: []string{
			"Площадь комнаты 14м².",
		},
	},
	{
		text: "Площадь комнаты 14 кв.м.",
		sentences: []string{
			"Площадь комнаты 14 кв.м.",
		},
	},
	{
		text: "1°C соответствует 33.8°F.",
		sentences: []string{
			"1°C соответствует 33.8°F.",
		},
	},
	{
		text: "Сегодня 27 октября 2014 года.",
		sentences: []string{
			"Сегодня 27 октября 2014 года.",
		},
	},
	{
		text: "Эта машина стоит 150 000 дол.!",
		sentences: []string{
			"Эта машина стоит 150 000 дол.!",
		},
	},
	{
		text: "Эта машина стоит $150 000!",
		sentences: []string{
			"Эта машина стоит $150 000!",
		},
	},
	{
		text: "Вот номер моего телефона: +39045969798. Передавайте привет г-ну Шапочкину. До свидания.",
		sentences: []string{
			"Вот номер моего телефона: +39045969798.",
			"Передавайте привет г-ну Шапочкину.",
			"До свидания.",
		},
	},
	{
		text: "Постойте, разве можно указывать цены в у.е.!",
		sentences: []string{
			"Постойте, разве можно указывать цены в у.е.!",
		},
	},
	{
		text: "Едем на скорости 90 км/ч в сторону пгт. Брагиновка, о котором мы так много слышали по ТВ!",
		sentences: []string{
			"Едем на скорости 90 км/ч в сторону пгт. Брагиновка, о котором мы так много слышали по ТВ!",
		},
	},
	{
		text: "Д-р ветеринарных наук А. И. Семенов и пр. выступали на этом семинаре.",
		sentences: []string{
			"Д-р ветеринарных наук А. И. Семенов и пр. выступали на этом семинаре.",
		},
	},
	{
		text: "Уважаемый проф. Семенов! Просьба до 20.10 сдать отчет на кафедру.",
		sentences: []string{
			"Уважаемый проф. Семенов!",
			"Просьба до 20.10 сдать отчет на кафедру.",
		},
	},
	{
		text: "Первоначальная стоимость этого комплекта 30 долл., но сейчас действует скидка. Предъявите дисконтную карту, пожалуйста!",
		sentences: []string{
			"Первоначальная стоимость этого комплекта 30 долл., но сейчас действует скидка.",
			"Предъявите дисконтную карту, пожалуйста!",
		},
	},
	{
		text: "Виктор съел пол-лимона и ушел по-английски из дома на ул. 1 Мая.",
		sentences: []string{
			"Виктор съел пол-лимона и ушел по-английски из дома на ул. 1 Мая.",
		},
	},
	{
		text: "Напоминаю Вам, что 25.10 день рождения у Маши К., нужно будет купить ей подарок.",
		sentences: []string{
			"Напоминаю Вам, что 25.10 день рождения у Маши К., нужно будет купить ей подарок.",
		},
	},
	{
		text: "В 2010-2012 гг. Виктор посещал г. Волгоград неоднократно.",
		sentences: []string{
			"В 2010-2012 гг. Виктор посещал г. Волгоград неоднократно.",
		},
	},
	{
		text: "Маленькая девочка бежала и кричала: «Не видали маму?»",
		sentences: []string{
			"Маленькая девочка бежала и кричала: «Не видали маму?»",
		},
	},
	{
		text: "Кв. 234 находится на 4 этаже.",
		sentences: []string{
			"Кв. 234 находится на 4 этаже.",
		},
	},
	{
		text: "Нужно купить 1)рыбу 2)соль.",
		sentences: []string{
			"Нужно купить 1)рыбу 2)соль.",
		},
	},
	{
		text: "Л.Н. Толстой написал \"Войну и мир\". Кроме Волконских, Л. Н. Толстой состоял в близком родстве с некоторыми другими аристократическими родами. Дом, где родился Л.Н.Толстой, 1898 г. В 1854 году дом продан по распоряжению писателя на вывоз в село Долгое.",
		sentences: []string{
			"Л.Н. Толстой написал \"Войну и мир\".",
			"Кроме Волконских, Л. Н. Толстой состоял в близком родстве с некоторыми другими аристократическими родами.",
			"Дом, где родился Л.Н.Толстой, 1898 г. В 1854 году дом продан по распоряжению писателя на вывоз в село Долгое.",
		},
	},
}

func TestRussian(t *testing.T) {
	factory := languages.LanguageFactory{}
	language := factory.CreateLanguage("ru")
	for _, tt := range rutests {
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
