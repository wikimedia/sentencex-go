package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/wikimedia/sentencex-go/languages"
)

type SegmentationTest struct {
	text      string
	sentences []string
}

var GOLDEN_EN_RULES = []SegmentationTest{
	//  1) Simple period to end sentence
	{
		text:      "Hello World. My name is Jonas.",
		sentences: []string{"Hello World.", "My name is Jonas."},
	},
	//  2) Question mark to end sentence
	{
		text: "What is your name? My name is Jonas.", sentences: []string{"What is your name?", "My name is Jonas."}},
	//  3) Exclamation point to end sentence
	{
		text: "There it is! I found it.", sentences: []string{"There it is!", "I found it."}},
	//  4) One letter upper case abbreviations
	{
		text: "My name is Jonas E. Smith.", sentences: []string{"My name is Jonas E. Smith."}},
	//  5) One letter lower case abbreviations
	{
		text: "Please turn to p. 55.", sentences: []string{"Please turn to p. 55."}},
	//  6) Two letter lower case abbreviations in the middle of a sentence
	{
		text: "Were Jane and co. at the party?", sentences: []string{"Were Jane and co. at the party?"}},
	//  7) Two letter upper case abbreviations in the middle of a sentence
	{
		text:      "They closed the deal with Pitt, Briggs & Co. at noon.",
		sentences: []string{"They closed the deal with Pitt, Briggs & Co. at noon."},
	},
	//  8) Two letter lower case abbreviations at the end of a sentence
	{
		text: "Let's ask Jane and co. They should know.", sentences: []string{"Let's ask Jane and co.", "They should know."}},
	//  9) Two letter upper case abbreviations at the end of a sentence
	{
		text:      "They closed the deal with Pitt, Briggs & Co. It closed yesterday.",
		sentences: []string{"They closed the deal with Pitt, Briggs & Co.", "It closed yesterday."},
	},
	//  10) Two letter (prepositive) abbreviations
	{
		text: "I can see Mt. Fuji from here.", sentences: []string{"I can see Mt. Fuji from here."}},
	//  11) Two letter (prepositive & postpositive) abbreviations
	{
		text:      "St. Michael's Church is on 5th st. near the light.",
		sentences: []string{"St. Michael's Church is on 5th st. near the light."},
	},
	//  12) Possesive two letter abbreviations
	{
		text: "That is JFK Jr.'s book.", sentences: []string{"That is JFK Jr.'s book."}},
	//  13) Multi-period abbreviations in the middle of a sentence
	{
		text:      "I visited the U.S.A. last year.",
		sentences: []string{"I visited the U.S.A. last year."}},
	//  14) Multi-period abbreviations at the end of a sentence
	{
		text:      "I live in the E.U. How about you?",
		sentences: []string{"I live in the E.U.", "How about you?"},
	},
	//  15) U.S. as sentence boundary
	{
		text:      "I live in the U.S. How about you?",
		sentences: []string{"I live in the U.S.", "How about you?"},
	},
	//  16) U.S. as non sentence boundary with next word capitalized
	{
		text:      "I work for the U.S. Government in Virginia.",
		sentences: []string{"I work for the U.S. Government in Virginia."},
	},
	//  17) U.S. as non sentence boundary
	{
		text: "I have lived in the U.S. for 20 years.", sentences: []string{"I have lived in the U.S. for 20 years."}},
	//  Most difficult sentence to crack
	//  18) A.M. / P.M. as non sentence boundary and sentence boundary
	{
		text: "At 5 a.m. Mr. Smith went to the bank. He left the bank at 6 P.M. Mr. Smith then went to the store.",
		sentences: []string{
			"At 5 a.m. Mr. Smith went to the bank.",
			"He left the bank at 6 P.M.",
			"Mr. Smith then went to the store.",
		},
	},
	//  19) Number as non sentence boundary
	{
		text: "She has $100.00 in her bag.", sentences: []string{"She has $100.00 in her bag."}},
	//  20) Number as sentence boundary
	{
		text: "She has $100.00. It is in her bag.", sentences: []string{"She has $100.00.", "It is in her bag."}},
	//  21) Parenthetical inside sentence
	{
		text: "He teaches science (He previously worked for 5 years as an engineer.) at the local University.",
		sentences: []string{
			"He teaches science (He previously worked for 5 years as an engineer.) at the local University.",
		},
	},
	//  22) Email addresses
	{
		text:      "Her email is Jane.Doe@example.com. I sent her an email.",
		sentences: []string{"Her email is Jane.Doe@example.com.", "I sent her an email."},
	},
	//  23) Web addresses
	{
		text: "The site is: https://www.example.50.com/new-site/awesome_content.html. Please check it out.",
		sentences: []string{
			"The site is: https://www.example.50.com/new-site/awesome_content.html.",
			"Please check it out.",
		},
	},
	//  24) Single quotations inside sentence
	{
		text:      "She turned to him, \"This is great.\" she said.",
		sentences: []string{"She turned to him, \"This is great.\" she said."},
	},
	//  25) Double quotations inside sentence
	{
		text:      "She turned to him, \"This is great.\" she said.",
		sentences: []string{"She turned to him, \"This is great.\" she said."},
	},
	//  26) Double quotations at the end of a sentence
	{
		text:      "She turned to him, \"This is great.\" She held the book out to show him.",
		sentences: []string{"She turned to him, \"This is great.\"", "She held the book out to show him."},
	},
	//  27) Double punctuation (exclamation point)
	{
		text: "Hello!! Long time no see.", sentences: []string{"Hello!!", "Long time no see."}},
	//  28) Double punctuation (question mark)
	{
		text: "Hello?? Who is there?", sentences: []string{"Hello??", "Who is there?"}},
	//  29) Double punctuation (exclamation point / question mark)
	{
		text: "Hello!? Is that you?", sentences: []string{"Hello!?", "Is that you?"}},
	//  30) Double punctuation (question mark / exclamation point)
	{
		text: "Hello?! Is that you?", sentences: []string{"Hello?!", "Is that you?"}},
	//  31) List (period followed by parens and no period to end item)
	//  (
	//      "1.) The first item 2.) The second item",
	//      sentences: []string{"1.) The first item", "2.) The second item"},
	//  ),
	//  //  32) List (period followed by parens and period to end item)
	//  (
	//      "1.) The first item. 2.) The second item.",
	//      sentences: []string{"1.) The first item.", "2.) The second item."},
	//  ),
	//  //  33) List (parens and no period to end item)
	//  (
	//      "1) The first item 2) The second item",
	//      sentences: []string{"1) The first item", "2) The second item"},
	//  ),
	//  //  34) List (parens and period to end item)
	{
		text:      "1) The first item. 2) The second item.",
		sentences: []string{"1) The first item.", "2) The second item."},
	},
	//  //  35) List (period to mark list and no period to end item)
	//  (
	//      "1. The first item 2. The second item",
	//      sentences: []string{"1. The first item", "2. The second item"},
	//  ),
	//  //  36) List (period to mark list and period to end item)
	//  (
	//      "1. The first item. 2. The second item.",
	//      sentences: []string{"1. The first item.", "2. The second item."},
	//  ),
	//  //  37) List with bullet
	//  (
	//      "• 9. The first item • 10. The second item",
	//      sentences: []string{"• 9. The first item", "• 10. The second item"},
	//  ),
	//  //  38) List with hypthen
	//  (
	//      "⁃9. The first item ⁃10. The second item",
	//      sentences: []string{"⁃9. The first item", "⁃10. The second item"},
	//  ),
	//  //  39) Alphabetical list
	//  (
	//      "a. The first item b. The second item c. The third list item",
	//      sentences: []string{"a. The first item", "b. The second item", "c. The third list item"},
	//  ),
	//  40) Geo Coordinates
	{
		text:      "You can find it at N°. 1026.253.553. That is where the treasure is.",
		sentences: []string{"You can find it at N°. 1026.253.553.", "That is where the treasure is."},
	},
	//  41) Named entities with an exclamation point
	{
		text:      "She works at Yahoo! in the accounting department.",
		sentences: []string{"She works at Yahoo! in the accounting department."},
	},
	//  42) I as a sentence boundary and I as an abbreviation
	{
		text:      "We make a good team, you and I. Did you see Albert I. Jones yesterday?",
		sentences: []string{"We make a good team, you and I.", "Did you see Albert I. Jones yesterday?"},
	},
	//  43) Ellipsis at end of quotation
	{
		text: "Thoreau argues that by simplifying one’s life, “the laws of the universe will appear less complex. . . .”",
		sentences: []string{
			"Thoreau argues that by simplifying one’s life, “the laws of the universe will appear less complex. . . .”",
		},
	},
	//  44) Ellipsis with square brackets
	{
		text:      "Bohr sentences: [...] used the analogy of parallel stairways sentences: [...]\" (Smith 55).",
		sentences: []string{"Bohr sentences: [...] used the analogy of parallel stairways sentences: [...]\" (Smith 55)."},
	},
	//  45) Ellipsis as sentence boundary (standard ellipsis rules)
	{
		text: "If words are left off at the end of a sentence, and that is all that is omitted, indicate the omission with ellipsis marks (preceded and followed by a space) and then indicate the end of the sentence with a period . . . . Next sentence.",
		sentences: []string{
			"If words are left off at the end of a sentence, and that is all that is omitted, indicate the omission with ellipsis marks (preceded and followed by a space) and then indicate the end of the sentence with a period . . . .",
			"Next sentence.",
		},
	},
	//  46) Ellipsis as sentence boundary (non-standard ellipsis rules)
	{
		text:      "I never meant that.... She left the store.",
		sentences: []string{"I never meant that....", "She left the store."},
	},
	//  47) Ellipsis as non sentence boundary
	{
		text: "I wasn’t really ... well, what I mean...see . . . what I'm saying, the thing is . . . I didn’t mean it.",
		sentences: []string{
			"I wasn’t really ... well, what I mean...see . . . what I'm saying, the thing is . . . I didn’t mean it.",
		},
	},
	//  48) 4-dot ellipsis
	{
		text: "One further habit which was somewhat weakened . . . was that of combining words into self-interpreting compounds. . . . The practice was not abandoned. . . .",
		sentences: []string{
			"One further habit which was somewhat weakened . . . was that of combining words into self-interpreting compounds.",
			". . . The practice was not abandoned. . . .",
		},
	},
}

func Benchmark() float64 {
	factory := languages.LanguageFactory{}
	language := factory.CreateLanguage("en")

	total_tests := len(GOLDEN_EN_RULES)
	pass_tests := 0
	fail_tests := 0
	for _, tt := range GOLDEN_EN_RULES {
		segmented := language.Segment(tt.text)
		if len(segmented) != len(tt.sentences) {
			fail_tests++
			continue
		} else {
			fail := false
			for i, actual_sentence := range segmented {
				if strings.TrimSpace(actual_sentence) != tt.sentences[i] && actual_sentence != tt.sentences[i] {
					fail = true
					break
				}
			}
			if fail {
				fail_tests++
				continue
			}
		}
		pass_tests++
	}
	// fmt.Printf("Total tests: %d, Passed tests: %d, Failed tests: %d\n", total_tests, pass_tests, fail_tests)
	// fmt.Printf("Total pass: %f \n", float64(pass_tests)/float64(total_tests) * 100)
	return float64(pass_tests) / float64(total_tests) * 100
}

func main() {
	pass_percentage := Benchmark()
	fmt.Printf("Pass percentage: %f \n", float64(pass_percentage))
	elapsed := time.Duration(0)
	// Run the benchmark 100 times and calculate average time taken
	for i := 0; i < 100; i++ {
		start := time.Now()
		Benchmark()
		elapsed += time.Since(start) // Unit is ns

	}
	fmt.Printf("Time taken(Avg over 100 runs): %d micro second \n", int64(elapsed/(100*1000)))
}
