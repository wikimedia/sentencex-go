package languages_test

import (
	"testing"
)

var bgtests = []SegmentationTest{
	{
		text:      "В първата половина на ноември т.г. ще бъде свикан Консултативният съвет за национална сигурност, обяви държавният глава.",
		sentences: []string{"В първата половина на ноември т.г. ще бъде свикан Консултативният съвет за национална сигурност, обяви държавният глава."},
	},
	{
		text:      "Компютърът е устройство с общо предназначение, което може да бъде програмирано да извършва набор от аритметични и/или логически операции. Възможността поредицата такива операции да бъде променяна позволява компютърът да се използва за решаването на теоретично всяка изчислителна/логическа задача. Обикновено целта на тези операции е обработката на въведена информация (данни), представена в цифров (дигитален) вид, резултатът от които може да се изведе в най-общо казано използваема форма.",
		sentences: []string{
			"Компютърът е устройство с общо предназначение, което може да бъде програмирано да извършва набор от аритметични и/или логически операции.",
			"Възможността поредицата такива операции да бъде променяна позволява компютърът да се използва за решаването на теоретично всяка изчислителна/логическа задача.",
			"Обикновено целта на тези операции е обработката на въведена информация (данни), представена в цифров (дигитален) вид, резултатът от които може да се изведе в най-общо казано използваема форма.",
		},
	},
	{
		text:      "Пл. \"20 Април\"",
		sentences: []string{"Пл. \"20 Април\""},
	},
	{
		text:      "Той поставя началото на могъща династия, която управлява в продължение на 150 г. Саргон надделява в двубой с владетеля на град Ур и разширява териториите на държавата си по долното течение на Тигър и Ефрат. Стойностни, вкл. български и руски",
		sentences: []string{
			"Той поставя началото на могъща династия, която управлява в продължение на 150 г. Саргон надделява в двубой с владетеля на град Ур и разширява териториите на държавата си по долното течение на Тигър и Ефрат.",
			"Стойностни, вкл. български и руски",
		},
	},
}

func TestBulgarian(t *testing.T) {
	LanguageTest(t, "bg",bgtests)
}