package languages_test

import (
	"testing"
)

var frtests = []SegmentationTest{
	{
		text: "Après avoir été l'un des acteurs du projet génome humain, le Genoscope met aujourd'hui le cap vers la génomique environnementale. L'exploitation des données de séquences, prolongée par l'identification expérimentale des fonctions biologiques, notamment dans le domaine de la biocatalyse, ouvrent des perspectives de développements en biotechnologie industrielle.",
		sentences: []string{
			"Après avoir été l'un des acteurs du projet génome humain, le Genoscope met aujourd'hui le cap vers la génomique environnementale.",
			"L'exploitation des données de séquences, prolongée par l'identification expérimentale des fonctions biologiques, notamment dans le domaine de la biocatalyse, ouvrent des perspectives de développements en biotechnologie industrielle.",
		},
	},
	{
		text: "\"Airbus livrera comme prévu 30 appareils 380 cette année avec en ligne de mire l'objectif d'équilibre financier du programme en 2015\", a-t-il ajouté.",
		sentences: []string{
			"\"Airbus livrera comme prévu 30 appareils 380 cette année avec en ligne de mire l'objectif d'équilibre financier du programme en 2015\", a-t-il ajouté.",
		},
	},
	{
		text: "À 11 heures ce matin, la direction ne décomptait que douze grévistes en tout sur la France : ce sont ceux du site de Saran (Loiret), dont l’effectif est de 809 salariés, dont la moitié d’intérimaires. Elle assure que ce mouvement « n’aura aucun impact sur les livraisons ».",
		sentences: []string{
			"À 11 heures ce matin, la direction ne décomptait que douze grévistes en tout sur la France : ce sont ceux du site de Saran (Loiret), dont l’effectif est de 809 salariés, dont la moitié d’intérimaires.",
			"Elle assure que ce mouvement « n’aura aucun impact sur les livraisons ».",
		},
	},
	{
		text: "Ce modèle permet d’afficher le texte « LL.AA.II.RR. » pour l’abréviation de « Leurs Altesses impériales et royales » avec son infobulle.",
		sentences: []string{
			"Ce modèle permet d’afficher le texte « LL.AA.II.RR. » pour l’abréviation de « Leurs Altesses impériales et royales » avec son infobulle.",
		},
	},
	{
		text: "Les derniers ouvrages de Intercept Ltd. sont ici.",
		sentences: []string{
			"Les derniers ouvrages de Intercept Ltd. sont ici.",
		},
	},
}

func TestFrench(t *testing.T) {
	LanguageTest(t, "fr", frtests)
}
