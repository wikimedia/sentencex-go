package languages_test

import (
	"strings"
	"testing"

	"github.com/wikimedia/sentencex-go/languages"
)

var estests = []struct {
	text      string
	sentences []string
	skip      bool
}{
	{
		text:      "¿Cómo está hoy? Espero que muy bien.",
		sentences: []string{"¿Cómo está hoy?", "Espero que muy bien."},
	},
	{
		text:      "¡Hola señorita! Espero que muy bien.",
		sentences: []string{"¡Hola señorita!", "Espero que muy bien."},
	},
	{
		text:      "Hola Srta. Ledesma. Buenos días, soy el Lic. Naser Pastoriza, y él es mi padre, el Dr. Naser.",
		sentences: []string{"Hola Srta. Ledesma.", "Buenos días, soy el Lic. Naser Pastoriza, y él es mi padre, el Dr. Naser."},
	},
	{
		text:      "¡La casa cuesta $170.500.000,00! ¡Muy costosa! Se prevé una disminución del 12.5% para el próximo año.",
		sentences: []string{"¡La casa cuesta $170.500.000,00!", "¡Muy costosa!", "Se prevé una disminución del 12.5% para el próximo año."},
	},
	{
		text:      "«Ninguna mente extraordinaria está exenta de un toque de demencia.», dijo Aristóteles.",
		sentences: []string{"«Ninguna mente extraordinaria está exenta de un toque de demencia.», dijo Aristóteles."},
	},
	{
		text:      "«Ninguna mente extraordinaria está exenta de un toque de demencia», dijo Aristóteles. Pablo, ¿adónde vas? ¡¿Qué viste?!",
		sentences: []string{"«Ninguna mente extraordinaria está exenta de un toque de demencia», dijo Aristóteles.", "Pablo, ¿adónde vas?", "¡¿Qué viste?!"},
	},
	{
		text:      "Admón. es administración o me equivoco.",
		sentences: []string{"Admón. es administración o me equivoco."},
	},
	{
		text:      "¡Hola Srta. Ledesma! ¿Cómo está hoy? Espero que muy bien.",
		sentences: []string{"¡Hola Srta. Ledesma!", "¿Cómo está hoy?", "Espero que muy bien."},
	},
	{
		text:      "Buenos días, soy el Lic. Naser Pastoriza, y él es mi padre, el Dr. Naser.",
		sentences: []string{"Buenos días, soy el Lic. Naser Pastoriza, y él es mi padre, el Dr. Naser."},
	},
	{
		text:      "He apuntado una cita para la siguiente fecha: Mar. 23 de Nov. de 2014. Gracias.",
		sentences: []string{"He apuntado una cita para la siguiente fecha: Mar. 23 de Nov. de 2014.", "Gracias."},
	},
	{
		text:      "Núm. de tel: 351.123.465.4. Envíe mis saludos a la Sra. Rescia.",
		sentences: []string{"Núm. de tel: 351.123.465.4.", "Envíe mis saludos a la Sra. Rescia."},
	},
	{
		text:      "Cero en la escala Celsius o de grados centígrados (0 °C) se define como el equivalente a 273.15 K, con una diferencia de temperatura de 1 °C equivalente a una diferencia de 1 Kelvin. Esto significa que 100 °C, definido como el punto de ebullición del agua, se define como el equivalente a 373.15 K.",
		sentences: []string{"Cero en la escala Celsius o de grados centígrados (0 °C) se define como el equivalente a 273.15 K, con una diferencia de temperatura de 1 °C equivalente a una diferencia de 1 Kelvin.", "Esto significa que 100 °C, definido como el punto de ebullición del agua, se define como el equivalente a 373.15 K."},
	},
	{
		text:      "Durante la primera misión del Discovery (30 Ago. 1984 15:08.10) tuvo lugar el lanzamiento de dos satélites de comunicación, el nombre de esta misión fue STS-41-D.",
		sentences: []string{"Durante la primera misión del Discovery (30 Ago. 1984 15:08.10) tuvo lugar el lanzamiento de dos satélites de comunicación, el nombre de esta misión fue STS-41-D."},
	},
	{
		text:      "Frase del gran José Hernández: \"Aquí me pongo a cantar / al compás de la vigüela, / que el hombre que lo desvela / una pena estrordinaria, / como la ave solitaria / con el cantar se consuela. / [...] \".",
		sentences: []string{"Frase del gran José Hernández: \"Aquí me pongo a cantar / al compás de la vigüela, / que el hombre que lo desvela / una pena estrordinaria, / como la ave solitaria / con el cantar se consuela. / [...] \"."},
	},
	{
		text:      "Citando a Criss Jami «Prefiero ser un artista a ser un líder, irónicamente, un líder tiene que seguir las reglas.», lo cual parece muy acertado.",
		sentences: []string{"Citando a Criss Jami «Prefiero ser un artista a ser un líder, irónicamente, un líder tiene que seguir las reglas.», lo cual parece muy acertado."},
	},
	{
		text:      "Cuando llegué, le estaba dando ejercicios a los niños, uno de los cuales era \"3 + (14/7).x = 5\". ¿Qué te parece?",
		sentences: []string{"Cuando llegué, le estaba dando ejercicios a los niños, uno de los cuales era \"3 + (14/7).x = 5\".", "¿Qué te parece?"},
	},
	{
		text:      "Se le pidió a los niños que leyeran los párrf. 5 y 6 del art. 4 de la constitución de los EE. UU..",
		sentences: []string{"Se le pidió a los niños que leyeran los párrf. 5 y 6 del art. 4 de la constitución de los EE. UU.."},
		skip:      true,
	},
	{
		text:      "Una de las preguntas realizadas en la evaluación del día Lun. 15 de Mar. fue la siguiente: \"Alumnos, ¿cuál es el resultado de la operación 1.1 + 4/5?\". Disponían de 1 min. para responder esa pregunta.",
		sentences: []string{"Una de las preguntas realizadas en la evaluación del día Lun. 15 de Mar. fue la siguiente: \"Alumnos, ¿cuál es el resultado de la operación 1.1 + 4/5?\".", "Disponían de 1 min. para responder esa pregunta."},
	},
	{
		text:      "La temperatura del motor alcanzó los 120.5°C. Afortunadamente, pudo llegar al final de carrera.",
		sentences: []string{"La temperatura del motor alcanzó los 120.5°C.", "Afortunadamente, pudo llegar al final de carrera."},
	},
	{
		text:      "El volumen del cuerpo es 3m³. ¿Cuál es la superficie de cada cara del prisma?",
		sentences: []string{"El volumen del cuerpo es 3m³.", "¿Cuál es la superficie de cada cara del prisma?"},
	},
	{
		text:      "La habitación tiene 20.55m². El living tiene 50.0m².",
		sentences: []string{"La habitación tiene 20.55m².", "El living tiene 50.0m²."},
	},
	{
		text:      "1°C corresponde a 33.8°F. ¿A cuánto corresponde 35°C?",
		sentences: []string{"1°C corresponde a 33.8°F.", "¿A cuánto corresponde 35°C?"},
	},
	{
		text:      "Hamilton ganó el último gran premio de Fórmula 1, luego de 1:39:02.619 Hs. de carrera, segundo resultó Massa, a una diferencia de 2.5 segundos. De esta manera se consagró ¡Campeón mundial!",
		sentences: []string{"Hamilton ganó el último gran premio de Fórmula 1, luego de 1:39:02.619 Hs. de carrera, segundo resultó Massa, a una diferencia de 2.5 segundos.", "De esta manera se consagró ¡Campeón mundial!"},
	},
	{
		text:      "¡La casa cuesta $170.500.000,00! ¡Muy costosa! Se prevé una disminución del 12.5% para el próximo año.",
		sentences: []string{"¡La casa cuesta $170.500.000,00!", "¡Muy costosa!", "Se prevé una disminución del 12.5% para el próximo año."},
	},
	{
		text:      "El corredor No. 103 arrivó 4°.",
		sentences: []string{"El corredor No. 103 arrivó 4°."},
	},
	{
		text:      "Hoy es 27/04/2014, y es mi cumpleaños. ¿Cuándo es el tuyo?",
		sentences: []string{"Hoy es 27/04/2014, y es mi cumpleaños.", "¿Cuándo es el tuyo?"},
	},
	{
		text:      "Aquí está la lista de compras para el almuerzo: 1.Helado, 2.Carne, 3.Arroz. ¿Cuánto costará? Quizás $12.5.",
		sentences: []string{"Aquí está la lista de compras para el almuerzo: 1.Helado, 2.Carne, 3.Arroz.", "¿Cuánto costará?", "Quizás $12.5."},
		skip:      true,
	},
	{
		text:      "1 + 1 es 2. 2 + 2 es 4. El auto es de color rojo.",
		sentences: []string{"1 + 1 es 2.", "2 + 2 es 4.", "El auto es de color rojo."},
	},
	{
		text:      "La máquina viajaba a 100 km/h. ¿En cuánto tiempo recorrió los 153 Km.?",
		sentences: []string{"La máquina viajaba a 100 km/h.", "¿En cuánto tiempo recorrió los 153 Km.?"},
	},
	{
		text:      "Explora oportunidades de carrera en el área de Salud en el Hospital de Northern en Mt. Kisco.",
		sentences: []string{"Explora oportunidades de carrera en el área de Salud en el Hospital de Northern en Mt. Kisco."},
	},
}

func TestSpanish(t *testing.T) {
	factory := languages.LanguageFactory{}
	language := factory.CreateLanguage("es")
	for _, tt := range estests {
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
