package linter

import (
	"io/ioutil"
	"log"
	"strings"
)

// WordSet describes a set of words
type WordSet map[string]bool

// NewWordSet creates a new case insensitive WordSet from a slice of strings
func NewWordSet(words []string) WordSet {
	ws := make(WordSet)

	for _, word := range words {
		ws[strings.ToLower(word)] = true
	}

	return ws
}

// Has returns a boolean indicating if the token value is present in the WordSet
func (w WordSet) Has(t *Token) bool {
	has, _ := w[strings.ToLower(t.Value)]
	return has
}

// loadDictionary loads a file containing a dictionary
func loadDictionary() []string {
	var dictionary []string

	file, err := ioutil.ReadFile("/usr/share/dict/words")
	if err != nil {
		log.Fatal(err)
	}

	dictionary = strings.Split(string(file), "\n")

	return dictionary
}

// IrregularWords is a collection of English irregular words
var IrregularWords = NewWordSet([]string{"awoken", "been", "born", "beat", "become", "begun", "bent", "beset", "bet", "bid", "bidden", "bound", "bitten", "bled", "blown", "broken", "bred", "brought", "broadcast", "built", "burnt", "burst", "bought", "cast", "caught", "chosen", "clung", "come", "cost", "crept", "cut", "dealt", "dug", "dived", "done", "drawn", "dreamt", "driven", "drunk", "eaten", "fallen", "fed", "felt", "fought", "found", "fit", "fled", "flung", "flown", "forbidden", "forgotten", "foregone", "forgiven", "forsaken", "frozen", "gotten", "given", "gone", "ground", "grown", "hung", "heard", "hidden", "hit", "held", "hurt", "kept", "knelt", "knit", "known", "laid", "led", "leapt", "learnt", "left", "lent", "let", "lain", "lighted", "lost", "made", "meant", "met", "misspelt", "mistaken", "mown", "overcome", "overdone", "overtaken", "overthrown", "paid", "pled", "proven", "put", "quit", "read", "rid", "ridden", "rung", "risen", "run", "sawn", "said", "seen", "sought", "sold", "sent", "set", "sewn", "shaken", "shaven", "shorn", "shed", "shone", "shod", "shot", "shown", "shrunk", "shut", "sung", "sunk", "sat", "slept", "slain", "slid", "slung", "slit", "smitten", "sown", "spoken", "sped", "spent", "spilt", "spun", "spit", "split", "spread", "sprung", "stood", "stolen", "stuck", "stung", "stunk", "stridden", "struck", "strung", "striven", "sworn", "swept", "swollen", "swum", "swung", "taken", "taught", "torn", "told", "thought", "thrived", "thrown", "thrust", "trodden", "understood", "upheld", "upset", "woken", "worn", "woven", "wed", "wept", "wound", "won", "withheld", "withstood", "wrung", "written"})

// LinkingVerbs is a collection of English forms of the verb 'be' also known as 'linking verb'
var LinkingVerbs = NewWordSet([]string{"am", "are", "were", "being", "is", "been", "was", "be"})

// WeaselWords is a collection of English weasel words
var WeaselWords = NewWordSet([]string{"many", "various", "very", "fairly", "several", "extremely", "exceedingly", "quite", "remarkably", "few", "surprisingly", "mostly", "largely", "huge", "tiny", "excellent", "interestingly", "significantly", "substantially", "clearly", "vast", "relatively", "completely"})

//DictionaryWords is a colection of English words
var DictionaryWords = NewWordSet(loadDictionary())
