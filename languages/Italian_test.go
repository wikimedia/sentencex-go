package languages_test

import "testing"

var ittests = []struct {
	text      string
	sentences []string
	skip      bool
}{
	{
		text:      "Salve Sig.ra Mengoni! Come sta oggi?",
		sentences: []string{"Salve Sig.ra Mengoni!", "Come sta oggi?"},
	},
	{
		text:      "Una lettera si può iniziare in questo modo «Il/la sottoscritto/a.».",
		sentences: []string{"Una lettera si può iniziare in questo modo «Il/la sottoscritto/a.»."},
	},
	{
		text:      "La casa costa 170.500.000,00€!",
		sentences: []string{"La casa costa 170.500.000,00€!"},
	},
	{
		text:      "Buongiorno! Sono l'Ing. Mengozzi. È presente l'Avv. Cassioni?",
		sentences: []string{"Buongiorno!", "Sono l'Ing. Mengozzi.", "È presente l'Avv. Cassioni?"},
	},
	{
		text:      "Mi fissi un appuntamento per mar. 23 Nov.. Grazie.",
		sentences: []string{"Mi fissi un appuntamento per mar. 23 Nov..", "Grazie."},
		skip:      true,
	},
	{
		text:      "Ecco il mio tel.:01234567. Mi saluti la Sig.na Manelli. Arrivederci.",
		sentences: []string{"Ecco il mio tel.:01234567.", "Mi saluti la Sig.na Manelli.", "Arrivederci."},
	},
	{
		text:      "La centrale meteor. si è guastata. Gli idraul. son dovuti andare a sistemarla.",
		sentences: []string{"La centrale meteor. si è guastata.", "Gli idraul. son dovuti andare a sistemarla."},
	},
	{
		text:      "Hanno creato un algoritmo allo st. d. arte. Si ringrazia lo psicol. Serenti.",
		sentences: []string{"Hanno creato un algoritmo allo st. d. arte.", "Si ringrazia lo psicol. Serenti."},
	},
	{
		text:      "Chiamate il V.Cte. delle F.P., adesso!",
		sentences: []string{"Chiamate il V.Cte. delle F.P., adesso!"},
	},
	{
		text:      "Giancarlo ha sostenuto l'esame di econ. az..",
		sentences: []string{"Giancarlo ha sostenuto l'esame di econ. az.."},
	},
	{
		text:      "Stava viaggiando a 90 km/h verso la provincia di TR quando il Dott. Mesini ha sentito un rumore e si fermò!",
		sentences: []string{"Stava viaggiando a 90 km/h verso la provincia di TR quando il Dott. Mesini ha sentito un rumore e si fermò!"},
	},
	{
		text:      "Egregio Dir. Amm., le faccio sapere che l'ascensore non funziona.",
		sentences: []string{"Egregio Dir. Amm., le faccio sapere che l'ascensore non funziona."},
	},
	{
		text:      "Stava mangiando e/o dormendo.",
		sentences: []string{"Stava mangiando e/o dormendo."},
	},
	{
		text:      "Ricordatevi che dom 25 Set. sarà il compleanno di Maria; dovremo darle un regalo.",
		sentences: []string{"Ricordatevi che dom 25 Set. sarà il compleanno di Maria; dovremo darle un regalo."},
	},
	{
		text:      "La politica è quella della austerità; quindi verranno fatti tagli agli sprechi.",
		sentences: []string{"La politica è quella della austerità; quindi verranno fatti tagli agli sprechi."},
	},
	{
		text:      "Nel tribunale, l'Avv. Fabrizi ha urlato \"Io, l'illustrissimo Fabrizi, vi si oppone!\".",
		sentences: []string{"Nel tribunale, l'Avv. Fabrizi ha urlato \"Io, l'illustrissimo Fabrizi, vi si oppone!\"."},
	},
	{
		text:      "Le parti fisiche di un computer (ad es. RAM, CPU, tastiera, mouse, etc.) sono definiti HW.",
		sentences: []string{"Le parti fisiche di un computer (ad es. RAM, CPU, tastiera, mouse, etc.) sono definiti HW."},
	},
	{
		text:      "La parola \"casa\" è sinonimo di abitazione.",
		sentences: []string{"La parola \"casa\" è sinonimo di abitazione."},
	},
	{
		text:      "La \"Mulino Bianco\" fa alimentari pre-confezionati.",
		sentences: []string{"La \"Mulino Bianco\" fa alimentari pre-confezionati."},
	},
	{
		text:      "\"Ei fu. Siccome immobile / dato il mortal sospiro / stette la spoglia immemore / orba di tanto spiro / [...]\" (Manzoni).",
		sentences: []string{"\"Ei fu. Siccome immobile / dato il mortal sospiro / stette la spoglia immemore / orba di tanto spiro / [...]\" (Manzoni)."},
	},
	{
		text:      "Una lettera si può iniziare in questo modo «Il/la sottoscritto/a ... nato/a a ...».",
		sentences: []string{"Una lettera si può iniziare in questo modo «Il/la sottoscritto/a ... nato/a a ...»."},
	},
	{
		text:      "Per casa, in uno degli esercizi per i bambini c\"era \"3 + (14/7) = 5\"",
		sentences: []string{"Per casa, in uno degli esercizi per i bambini c\"era \"3 + (14/7) = 5\""},
	},
	{
		text:      "Ai bambini è stato chiesto di fare \"4:2*2\"",
		sentences: []string{"Ai bambini è stato chiesto di fare \"4:2*2\""},
	},
	{
		text:      "La maestra esclamò: \"Bambini, quanto fa '2/3 + 4/3?'\".",
		sentences: []string{"La maestra esclamò: \"Bambini, quanto fa '2/3 + 4/3?'\"."},
	},
	{
		text:      "Il motore misurava 120°C.",
		sentences: []string{"Il motore misurava 120°C."},
	},
	{
		text:      "Il volume era di 3m³.",
		sentences: []string{"Il volume era di 3m³."},
	},
	{
		text:      "La stanza misurava 20m².",
		sentences: []string{"La stanza misurava 20m²."},
	},
	{
		text:      "1°C corrisponde a 33.8°F.",
		sentences: []string{"1°C corrisponde a 33.8°F."},
	},
	{
		text:      "Oggi è il 27-10-14.",
		sentences: []string{"Oggi è il 27-10-14."},
	},
	{
		text:      "La casa costa 170.500.000,00€!",
		sentences: []string{"La casa costa 170.500.000,00€!"},
	},
	{
		text:      "Il corridore 103 è arrivato 4°.",
		sentences: []string{"Il corridore 103 è arrivato 4°."},
	},
	{
		text:      "Oggi è il 27/10/2014.",
		sentences: []string{"Oggi è il 27/10/2014."},
	},
	{
		text:      "Ecco l'elenco: 1.gelato, 2.carne, 3.riso.",
		sentences: []string{"Ecco l'elenco: 1.gelato, 2.carne, 3.riso."},
	},
	{
		text:      "Devi comprare : 1)pesce 2)sale.",
		sentences: []string{"Devi comprare : 1)pesce 2)sale."},
	},
	{
		text:      "La macchina viaggiava a 100 km/h.",
		sentences: []string{"La macchina viaggiava a 100 km/h."},
	},
}


func TestItalian(t *testing.T) {
	LanguageTest(t, "it", ittests)
}
