package languages_test

import (
	"strings"
	"testing"

	"github.com/wikimedia/sentencex-go/languages"
)

var tests = []struct {
	text      string
	sentences []string
	skip      bool
}{
	{
		text:      "This is Dr. Watson",
		sentences: []string{"This is Dr. Watson"}},
	{
		text:      "Roses Are Red. Violets Are Blue",
		sentences: []string{"Roses Are Red.", "Violets Are Blue"},
	},
	{
		text:      "Hello! How are you?",
		sentences: []string{"Hello!", "How are you?"}},
	{
		text:      "This is a test.",
		sentences: []string{"This is a test."},
	},
	{
		text:      "Mr. Smith went to Washington.",
		sentences: []string{"Mr. Smith went to Washington."},
	},
	{text: "What a suprise?!", sentences: []string{"What a suprise?!"}},
	{text: "That's all folks...", sentences: []string{"That's all folks..."}},
	{
		text:      "First line\nSecond line",
		sentences: []string{"First line\nSecond line"},
	},
	{
		text:      "First line\nSecond line\n\nThird line",
		sentences: []string{"First line\nSecond line", "\n\n", "Third line"},
	},
	{
		text: "This is UK. Not US", sentences: []string{"This is UK.", "Not US"}},
	{
		text:      "This balloon costs $1.20",
		sentences: []string{"This balloon costs $1.20"}},
	{
		text:      "Hello World. My name is Jonas.",
		sentences: []string{"Hello World.", "My name is Jonas."},
	},
	{text: "What is your name? My name is Jonas.", sentences: []string{"What is your name?", "My name is Jonas."}},
	{text: "There it is! I found it.", sentences: []string{"There it is!", "I found it."}},
	{text: "My name is Jonas E. Smith.", sentences: []string{"My name is Jonas E. Smith."}},
	{text: "Please turn to p. 55.", sentences: []string{"Please turn to p. 55."}},
	{text: "Were Jane and co. at the party?", sentences: []string{"Were Jane and co. at the party?"}},
	{text: "They closed the deal with Pitt, Briggs & Co. at noon.", sentences: []string{"They closed the deal with Pitt, Briggs & Co. at noon."}},
	{
		text:      "Let's ask Jane and co. They should know.",
		sentences: []string{"Let's ask Jane and co.", "They should know."},
		// Acceptable:
		// ["Let's ask Jane and co. They should know."],
		skip: true,
	},
	{
		text: "They closed the deal with Pitt, Briggs & Co. It closed yesterday.", sentences: []string{"They closed the deal with Pitt, Briggs & Co.", "It closed yesterday."},
		// Acceptable:
		// ["They closed the deal with Pitt, Briggs & Co. It closed yesterday."],
		skip: true,
	},
	{text: "I can see Mt. Fuji from here.", sentences: []string{"I can see Mt. Fuji from here."}},
	{text: "St. Michael's Church is on 5th st. near the light.", sentences: []string{"St. Michael's Church is on 5th st. near the light."}},
	{text: "That is JFK Jr.'s book.", sentences: []string{"That is JFK Jr.'s book."}},
	{text: "I visited the U.S.A. last year.", sentences: []string{"I visited the U.S.A. last year."}},
	{
		text: "I live in the E.U. How about you?", sentences: []string{"I live in the E.U.", "How about you?"},
		skip: true,
	},
	{
		text:      "I live in the U.S. How about you?",
		sentences: []string{"I live in the U.S.", "How about you?"},
		skip:      true,
	},
	{text: "I work for the U.S. Government in Virginia.", sentences: []string{"I work for the U.S. Government in Virginia."}},
	{text: "I have lived in the U.S. for 20 years.", sentences: []string{"I have lived in the U.S. for 20 years."}},
	{
		text:      "At 5 a.m. Mr. Smith went to the bank. He left the bank at 6 P.M. Mr. Smith then went to the store.",
		sentences: []string{"At 5 a.m. Mr. Smith went to the bank.", "He left the bank at 6 P.M.", "Mr. Smith then went to the store."},
		skip:      true,
	},
	{text: "She has $100.00 in her bag.", sentences: []string{"She has $100.00 in her bag."}},
	{text: "She has $100.00. It is in her bag.", sentences: []string{"She has $100.00.", "It is in her bag."}},
	{text: "He teaches science (He previously worked for 5 years as an engineer.) at the local University.", sentences: []string{"He teaches science (He previously worked for 5 years as an engineer.) at the local University."}},
	{text: "Her email is Jane.Doe@example.com. I sent her an email.", sentences: []string{"Her email is Jane.Doe@example.com.", "I sent her an email."}},
	{text: "The site is, https,//www.example.50.com/new-site/awesome_content.html. Please check it out.", sentences: []string{"The site is, https,//www.example.50.com/new-site/awesome_content.html.", "Please check it out."}},
	{text: "She turned to him, 'This is great.' she said.", sentences: []string{"She turned to him, 'This is great.' she said."}},
	{text: "She turned to him, \"This is great.\" she said.", sentences: []string{"She turned to him, \"This is great.\" she said."}},
	{text: "She turned to him, \"This is great.\" She held the book out to show him.", sentences: []string{"She turned to him, \"This is great.\"", "She held the book out to show him."}, skip: true},
	{text: "Hello!! Long time no see.", sentences: []string{"Hello!!", "Long time no see."}},
	{text: "Hello?? Who is there?", sentences: []string{"Hello??", "Who is there?"}},
	{text: "Hello!? Is that you?", sentences: []string{"Hello!?", "Is that you?"}},
	{text: "Hello?! Is that you?", sentences: []string{"Hello?!", "Is that you?"}},
	{text: "You can find it at N°. 1026.253.553. That is where the treasure is.", sentences: []string{"You can find it at N°. 1026.253.553.", "That is where the treasure is."}},
	{text: "She works at Yahoo! in the accounting department.", sentences: []string{"She works at Yahoo! in the accounting department."}},
	{text: "We make a good team, you and I. Did you see Albert I. Jones yesterday?", sentences: []string{"We make a good team, you and I.", "Did you see Albert I. Jones yesterday?"}, skip: true},
	{text: "Thoreau argues that by simplifying one’s life, “the laws of the universe will appear less complex. . . .”", sentences: []string{"Thoreau argues that by simplifying one’s life, “the laws of the universe will appear less complex. . . .”"}},
	{text: "\"Bohr [...] used the analogy of parallel stairways [...]\" (Smith 55).", sentences: []string{"\"Bohr [...] used the analogy of parallel stairways [...]\" (Smith 55)."}},
	{text: "If words are left off at the end of a sentence, and that is all that is omitted, indicate the omission with ellipsis marks (preceded and followed by a space) and then indicate the end of the sentence with a period . . . . Next sentence.", sentences: []string{"If words are left off at the end of a sentence, and that is all that is omitted, indicate the omission with ellipsis marks (preceded and followed by a space) and then indicate the end of the sentence with a period . . . .", "Next sentence."}, skip: true},
	{text: "I never meant that.... She left the store.", sentences: []string{"I never meant that....", "She left the store."}},
	{text: "I wasn’t really ... well, what I mean...see . . . what I'm saying, the thing is . . . I didn’t mean it.", sentences: []string{"I wasn’t really ... well, what I mean...see . . . what I'm saying, the thing is . . . I didn’t mean it."}, skip: true},
	{text: "One further habit which was somewhat weakened . . . was that of combining words into self-interpreting compounds. . . . The practice was not abandoned. . . .", sentences: []string{"One further habit which was somewhat weakened . . . was that of combining words into self-interpreting compounds.", ". . . The practice was not abandoned. . . ."}, skip: true},
	{text: "Saint Maximus (died 250) is a Christian saint and martyr.[1] The emperor Decius published a decree ordering the veneration of busts of the deified emperors.", sentences: []string{"Saint Maximus (died 250) is a Christian saint and martyr.[1]", "The emperor Decius published a decree ordering the veneration of busts of the deified emperors."}},
	{text: "Differing agendas can potentially create an understanding gap in a consultation.11 12 Take the example of one of the most common presentations in ill health: the common cold.", sentences: []string{"Differing agendas can potentially create an understanding gap in a consultation.11 12 Take the example of one of the most common presentations in ill health: the common cold."}},
	{text: "Its traditional use[1] is well documented in the ethnobotanical literature [2–11]. Leaves, buds, tar and essential oils are used to treat a wide spectrum of diseases.", sentences: []string{"Its traditional use[1] is well documented in the ethnobotanical literature [2–11].", "Leaves, buds, tar and essential oils are used to treat a wide spectrum of diseases."}},
	{text: "Thus increasing the desire for political reform both in Lancashire and in the country at large.[7][8] This was a serious misdemeanour,[16] encouraging them to declare the assembly illegal as soon as it was announced on 31 July.[17][18] The radicals sought a second opinion on the meeting's legality.", sentences: []string{"Thus increasing the desire for political reform both in Lancashire and in the country at large.[7][8]", "This was a serious misdemeanour,[16] encouraging them to declare the assembly illegal as soon as it was announced on 31 July.[17][18]", "The radicals sought a second opinion on the meeting's legality."}},
	{
		text:      "“Why, indeed?” murmured Holmes. “Your Majesty had not spoken before I \nwas aware that I was addressing Wilhelm Gottsreich Sigismond von \nOrmstein, Grand Duke of Cassel-Felstein, and hereditary King of \nBohemia.”",
		sentences: []string{"“Why, indeed?” murmured Holmes.", "“Your Majesty had not spoken before I \nwas aware that I was addressing Wilhelm Gottsreich Sigismond von \nOrmstein, Grand Duke of Cassel-Felstein, and hereditary King of \nBohemia.”"},
	},
	{
		text:      "“How many? I don’t know.”",
		sentences: []string{"“How many? I don’t know.”"},
	},
}

func TestEnglish(t *testing.T) {
	factory := languages.LanguageFactory{}
	english := factory.CreateLanguage("en")
	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			if tt.skip {
				t.Skip()
			}
			segmented := english.Segment(tt.text)
			if len(segmented) != len(tt.sentences) {
				t.Errorf("Expected %d sentences, got %d", len(tt.sentences), len(segmented))
				t.Error(segmented)
			} else {
				for i, actual_sentence := range segmented {
					if (strings.TrimSpace(actual_sentence) != tt.sentences[i]) && ( actual_sentence != tt.sentences[i]) {
						t.Errorf("Expected '%s', got '%s'", tt.sentences[i], actual_sentence)
					}
				}
			}
		})
	}
}
