package languages_test

import (
	"testing"
)

var pttests = []SegmentationTest{
	{
		text: "A Lei do Sorteio (n.º 1860, de 4 de janeiro de 1908) introduziu o serviço militar obrigatório para as Forças Armadas do Brasil, implantado de fato em 1916, substituindo o recrutamento forçado, o antigo “tributo de sangue”, e permitindo a constituição de uma reserva.",
		sentences: []string{
			"A Lei do Sorteio (n.º 1860, de 4 de janeiro de 1908) introduziu o serviço militar obrigatório para as Forças Armadas do Brasil, implantado de fato em 1916, substituindo o recrutamento forçado, o antigo “tributo de sangue”, e permitindo a constituição de uma reserva.",
		},
	},
	{
		text: "Os oficiais mantinham a disciplina pelo castigo corporal.[13] Na Marinha, isso resultou na Revolta da Chibata de 1910.[14]",
		sentences: []string{
			"Os oficiais mantinham a disciplina pelo castigo corporal.[13]",
			"Na Marinha, isso resultou na Revolta da Chibata de 1910.[14]",
		},
	},
	{
		text: "A nova legislação era a lei 2.556, de 26 de setembro de 1874, e o decreto 5.881, de 17 de fevereiro de 1875.[35]",
		sentences: []string{
			"A nova legislação era a lei 2.556, de 26 de setembro de 1874, e o decreto 5.881, de 17 de fevereiro de 1875.[35]",
		},
	},
}

func TestPortuguese(t *testing.T) {
	LanguageTest(t, "pt", pttests)
}
