package languages_test

import (
	"strings"
	"testing"

	"github.com/wikimedia/sentencex-go/languages"
)

var detests = []struct {
    text      string
    sentences []string
    skip      bool
}{
    {
        text:      "„Ich habe heute keine Zeit“, sagte die Frau und flüsterte leise: „Und auch keine Lust.“ Wir haben 1.000.000 Euro.",
        sentences: []string{"„Ich habe heute keine Zeit“, sagte die Frau und flüsterte leise: „Und auch keine Lust.“", "Wir haben 1.000.000 Euro."},
    },
    {
        text:      "Es gibt jedoch einige Vorsichtsmaßnahmen, die Du ergreifen kannst, z. B. ist es sehr empfehlenswert, dass Du Dein Zuhause von allem Junkfood befreist.",
        sentences: []string{"Es gibt jedoch einige Vorsichtsmaßnahmen, die Du ergreifen kannst, z. B. ist es sehr empfehlenswert, dass Du Dein Zuhause von allem Junkfood befreist."},
    },
    {
        text:      "Was sind die Konsequenzen der Abstimmung vom 12. Juni?",
        sentences: []string{"Was sind die Konsequenzen der Abstimmung vom 12. Juni?"},
    },
    {
        text:      "Thomas sagte: ,,Wann kommst zu mir?” ,,Das weiß ich noch nicht“, antwortete Susi, ,,wahrscheinlich am Sonntag.“ Wir haben 1.000.000 Euro.",
        sentences: []string{"Thomas sagte: ,,Wann kommst zu mir?” ,,Das weiß ich noch nicht“, antwortete Susi, ,,wahrscheinlich am Sonntag.“", "Wir haben 1.000.000 Euro."},
        skip:      true,
    },
    {
        text:      "„Lass uns jetzt essen gehen!“, sagte die Mutter zu ihrer Freundin, „am besten zum Italiener.“",
        sentences: []string{"„Lass uns jetzt essen gehen!“, sagte die Mutter zu ihrer Freundin, „am besten zum Italiener.“"},
    },
    {
        text:      "Wir haben 1.000.000 Euro.",
        sentences: []string{"Wir haben 1.000.000 Euro."},
    },
    {
        text:      "Sie bekommen 3,50 Euro zurück.",
        sentences: []string{"Sie bekommen 3,50 Euro zurück."},
    },
    {
        text:      "Dafür brauchen wir 5,5 Stunden.",
        sentences: []string{"Dafür brauchen wir 5,5 Stunden."},
    },
    {
        text:      "Bitte überweisen Sie 5.300,25 Euro.",
        sentences: []string{"Bitte überweisen Sie 5.300,25 Euro."},
    },
    {
        text:      "1. Dies ist eine Punkteliste.",
        sentences: []string{"1. Dies ist eine Punkteliste."},
        skip:      true,
    },
    {
        text:      "Wir trafen Dr. med. Meyer in der Stadt.",
        sentences: []string{"Wir trafen Dr. med. Meyer in der Stadt."},
    },
    {
        text:      "Wir brauchen Getränke, z. B. Wasser, Saft, Bier usw.",
        sentences: []string{"Wir brauchen Getränke, z. B. Wasser, Saft, Bier usw."},
    },
    {
        text:      "Ich kann u.a. Spanisch sprechen.",
        sentences: []string{"Ich kann u.a. Spanisch sprechen."},
    },
    {
        text:      "Frau Prof. Schulze ist z. Z. nicht da.",
        sentences: []string{"Frau Prof. Schulze ist z. Z. nicht da."},
    },
    {
        text:      "Sie erhalten ein neues Bank-Statement bzw. ein neues Schreiben.",
        sentences: []string{"Sie erhalten ein neues Bank-Statement bzw. ein neues Schreiben."},
    },
    {
        text:      "Z. T. ist die Lieferung unvollständig.",
        sentences: []string{"Z. T. ist die Lieferung unvollständig."},
    },
    {
        text:      "Das finden Sie auf S. 225.",
        sentences: []string{"Das finden Sie auf S. 225."},
    },
    {
        text:      "Sie besucht eine kath. Schule.",
        sentences: []string{"Sie besucht eine kath. Schule."},
    },
    {
        text:      "Wir benötigen Zeitungen, Zeitschriften u. Ä. für unser Projekt.",
        sentences: []string{"Wir benötigen Zeitungen, Zeitschriften u. Ä. für unser Projekt."},
    },
    {
        text:      "Das steht auf S. 23, s. vorherige Anmerkung.",
        sentences: []string{"Das steht auf S. 23, s. vorherige Anmerkung."},
    },
    {
        text:      "Dies ist meine Adresse: Dr. Meier, Berliner Str. 5, 21234 Bremen.",
        sentences: []string{"Dies ist meine Adresse: Dr. Meier, Berliner Str. 5, 21234 Bremen."},
    },
    {
        text:      "Er sagte: „Hallo, wie geht´s Ihnen, Frau Prof. Müller?“",
        sentences: []string{"Er sagte: „Hallo, wie geht´s Ihnen, Frau Prof. Müller?“"},
    },
    {
        text:      "Fit in vier Wochen\n\nDeine Anleitung für eine reine Ernährung und ein gesünderes und glücklicheres Leben\n\nRECHTLICHE HINWEISE\n\nOhne die ausdrückliche schriftliche Genehmigung der Eigentümerin von instafemmefitness, Anna Anderson, darf dieses E-Book weder teilweise noch in vollem Umfang reproduziert, gespeichert, kopiert oder auf irgendeine Weise übertragen werden. Wenn Du das E-Book auf einem öffentlich zugänglichen Computer ausdruckst, musst Du es nach dem Ausdrucken von dem Computer löschen. Jedes E-Book wird mit einem Benutzernamen und Transaktionsinformationen versehen.\n\nVerstöße gegen dieses Urheberrecht werden im vollen gesetzlichen Umfang geltend gemacht. Obgleich die Autorin und Herausgeberin alle Anstrengungen unternommen hat, sicherzustellen, dass die Informationen in diesem Buch zum Zeitpunkt der Drucklegung korrekt sind, übernimmt die Autorin und Herausgeberin keine Haftung für etwaige Verluste, Schäden oder Störungen, die durch Fehler oder Auslassungen in Folge von Fahrlässigkeit, zufälligen Umständen oder sonstigen Ursachen entstehen, und lehnt hiermit jedwede solche Haftung ab.\n\nDieses Buch ist kein Ersatz für die medizinische Beratung durch Ärzte. Der Leser/die Leserin sollte regelmäßig einen Arzt/eine Ärztin hinsichtlich Fragen zu seiner/ihrer Gesundheit und vor allem in Bezug auf Symptome, die eventuell einer ärztlichen Diagnose oder Behandlung bedürfen, konsultieren.\n\nDie Informationen in diesem Buch sind dazu gedacht, ein ordnungsgemäßes Training zu ergänzen, nicht aber zu ersetzen. Wie jeder andere Sport, der Geschwindigkeit, Ausrüstung, Gleichgewicht und Umweltfaktoren einbezieht, stellt dieser Sport ein gewisses Risiko dar. Die Autorin und Herausgeberin rät den Lesern dazu, die volle Verantwortung für die eigene Sicherheit zu übernehmen und die eigenen Grenzen zu beachten. Vor dem Ausüben der in diesem Buch beschriebenen Übungen solltest Du sicherstellen, dass Deine Ausrüstung in gutem Zustand ist, und Du solltest keine Risiken außerhalb Deines Erfahrungs- oder Trainingsniveaus, Deiner Fähigkeiten oder Deines Komfortbereichs eingehen.\nHintergrundillustrationen Urheberrecht © 2013 bei Shuttershock, Buchgestaltung und -produktion durch Anna Anderson Verfasst von Anna Anderson\nUrheberrecht © 2014 Instafemmefitness. Alle Rechte vorbehalten\n\nÜber mich",
        sentences: []string{
            "Fit in vier Wochen",
            "Deine Anleitung für eine reine Ernährung und ein gesünderes und glücklicheres Leben",
            "RECHTLICHE HINWEISE",
            "Ohne die ausdrückliche schriftliche Genehmigung der Eigentümerin von instafemmefitness, Anna Anderson, darf dieses E-Book weder teilweise noch in vollem Umfang reproduziert, gespeichert, kopiert oder auf irgendeine Weise übertragen werden.",
            "Wenn Du das E-Book auf einem öffentlich zugänglichen Computer ausdruckst, musst Du es nach dem Ausdrucken von dem Computer löschen.",
            "Jedes E-Book wird mit einem Benutzernamen und Transaktionsinformationen versehen.",
            "Verstöße gegen dieses Urheberrecht werden im vollen gesetzlichen Umfang geltend gemacht.",
            "Obgleich die Autorin und Herausgeberin alle Anstrengungen unternommen hat, sicherzustellen, dass die Informationen in diesem Buch zum Zeitpunkt der Drucklegung korrekt sind, übernimmt die Autorin und Herausgeberin keine Haftung für etwaige Verluste, Schäden oder Störungen, die durch Fehler oder Auslassungen in Folge von Fahrlässigkeit, zufälligen Umständen oder sonstigen Ursachen entstehen, und lehnt hiermit jedwede solche Haftung ab.",
            "Dieses Buch ist kein Ersatz für die medizinische Beratung durch Ärzte.",
            "Der Leser/die Leserin sollte regelmäßig einen Arzt/eine Ärztin hinsichtlich Fragen zu seiner/ihrer Gesundheit und vor allem in Bezug auf Symptome, die eventuell einer ärztlichen Diagnose oder Behandlung bedürfen, konsultieren.",
            "Die Informationen in diesem Buch sind dazu gedacht, ein ordnungsgemäßes Training zu ergänzen, nicht aber zu ersetzen.",
            "Wie jeder andere Sport, der Geschwindigkeit, Ausrüstung, Gleichgewicht und Umweltfaktoren einbezieht, stellt dieser Sport ein gewisses Risiko dar.",
            "Die Autorin und Herausgeberin rät den Lesern dazu, die volle Verantwortung für die eigene Sicherheit zu übernehmen und die eigenen Grenzen zu beachten.",
            "Vor dem Ausüben der in diesem Buch beschriebenen Übungen solltest Du sicherstellen, dass Deine Ausrüstung in gutem Zustand ist, und Du solltest keine Risiken außerhalb Deines Erfahrungs- oder Trainingsniveaus, Deiner Fähigkeiten oder Deines Komfortbereichs eingehen.",
            "Hintergrundillustrationen Urheberrecht © 2013 bei Shuttershock, Buchgestaltung und -produktion durch Anna Anderson Verfasst von Anna Anderson",
            "Urheberrecht © 2014 Instafemmefitness.",
            "Alle Rechte vorbehalten",
            "Über mich",
        },
        skip: true,
    },
    {
        text:      "Es gibt jedoch einige Vorsichtsmaßnahmen, die Du ergreifen kannst, z. B. ist es sehr empfehlenswert, dass Du Dein Zuhause von allem Junkfood befreist. Ich persönlich kaufe kein Junkfood oder etwas, das nicht rein ist (ich traue mir da selbst nicht!). Ich finde jeden Vorwand, um das Junkfood zu essen, vor allem die Vorstellung, dass ich nicht mehr in Versuchung kommen werde, wenn ich es jetzt aufesse und es weg ist. Es ist schon komisch, was unser Verstand mitunter anstellt!",
        sentences: []string{
            "Es gibt jedoch einige Vorsichtsmaßnahmen, die Du ergreifen kannst, z. B. ist es sehr empfehlenswert, dass Du Dein Zuhause von allem Junkfood befreist.",
            "Ich persönlich kaufe kein Junkfood oder etwas, das nicht rein ist (ich traue mir da selbst nicht!).",
            "Ich finde jeden Vorwand, um das Junkfood zu essen, vor allem die Vorstellung, dass ich nicht mehr in Versuchung kommen werde, wenn ich es jetzt aufesse und es weg ist.",
            "Es ist schon komisch, was unser Verstand mitunter anstellt!",
        },
        skip: true,
    },
    {
        text:      "Ob Sie in Hannover nur auf der Durchreise, für einen längeren Aufenthalt oder zum Besuch einer der zahlreichen Messen sind: Die Hauptstadt des Landes Niedersachsens hat viele Sehenswürdigkeiten und ist zu jeder Jahreszeit eine Reise Wert. Hannovers Ursprünge können bis zur römischen Kaiserzeit zurückverfolgt werden, und zwar durch Ausgrabungen von Tongefäßen aus dem 1. -3. Jahrhundert nach Christus, die an mehreren Stellen im Untergrund des Stadtzentrums durchgeführt wurden.",
        sentences: []string{
            "Ob Sie in Hannover nur auf der Durchreise, für einen längeren Aufenthalt oder zum Besuch einer der zahlreichen Messen sind: Die Hauptstadt des Landes Niedersachsens hat viele Sehenswürdigkeiten und ist zu jeder Jahreszeit eine Reise Wert.",
            "Hannovers Ursprünge können bis zur römischen Kaiserzeit zurückverfolgt werden, und zwar durch Ausgrabungen von Tongefäßen aus dem 1. -3. Jahrhundert nach Christus, die an mehreren Stellen im Untergrund des Stadtzentrums durchgeführt wurden.",
        },
        skip: true,
    },
    {
        text:      "• 3. Seien Sie achtsam bei der Auswahl der Nahrungsmittel! \n• 4. Nehmen Sie zusätzlich Folsäurepräparate und essen Sie Fisch! \n• 5. Treiben Sie regelmäßig Sport! \n• 6. Beginnen Sie mit Übungen für die Beckenbodenmuskulatur! \n• 7. Reduzieren Sie Ihren Alkoholgenuss! \n",
        sentences: []string{
            "• 3. Seien Sie achtsam bei der Auswahl der Nahrungsmittel!",
            "• 4. Nehmen Sie zusätzlich Folsäurepräparate und essen Sie Fisch!",
            "• 5. Treiben Sie regelmäßig Sport!",
            "• 6. Beginnen Sie mit Übungen für die Beckenbodenmuskulatur!",
            "• 7. Reduzieren Sie Ihren Alkoholgenuss!",
        },
        skip: true,
    },
    {
        text:      "Was pro Jahr10. Zudem pro Jahr um 0.3 %11. Der gängigen Theorie nach erfolgt der Anstieg.",
        sentences: []string{
            "Was pro Jahr10.",
            "Zudem pro Jahr um 0.3 %11.",
            "Der gängigen Theorie nach erfolgt der Anstieg.",
        },
    },
    {
        text:      "s. vorherige Anmerkung.",
        sentences: []string{"s. vorherige Anmerkung."},
    },
    {
        text:      "Mit Inkrafttreten des Mindestlohngesetzes (MiLoG) zum 01. Januar 2015 werden in Bezug auf den Einsatz von Leistungs.",
        sentences: []string{"Mit Inkrafttreten des Mindestlohngesetzes (MiLoG) zum 01. Januar 2015 werden in Bezug auf den Einsatz von Leistungs."},
    },
    {
        text:      "• einige Sorten Weichkäse  \n• rohes oder nicht ganz durchgebratenes Fleisch  \n• ungeputztes Gemüse und ungewaschener Salat  \n• nicht ganz durchgebratenes Hühnerfleisch, rohe oder nur weich gekochte Eier",
        sentences: []string{
            "• einige Sorten Weichkäse",
            "• rohes oder nicht ganz durchgebratenes Fleisch",
            "• ungeputztes Gemüse und ungewaschener Salat",
            "• nicht ganz durchgebratenes Hühnerfleisch, rohe oder nur weich gekochte Eier",
        },
        skip: true,
    },
}

func TestDeutsch(t *testing.T) {
	factory := languages.LanguageFactory{}
	language := factory.CreateLanguage("de")
	for _, tt := range detests {
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
