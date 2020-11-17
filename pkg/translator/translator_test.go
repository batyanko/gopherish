package translator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test the logic of our translator against the universally recognized English - Gopherish dictionary.

// Dictionary for words starting with a vowel
var vowelDict = map[string]string{
	"apple":     "gapple",
	"oomph":     "goomph",
	"interest":  "ginterest",
	"use":       "guse",
	"entertain": "gentertain",
	"yppee":     "gyppee",
	"you":       "gyou",

	"Apple":     "Gapple",
	"Oomph":     "Goomph",
	"Interest":  "Ginterest",
	"Use":       "Guse",
	"Entertain": "Gentertain",
	"Yppee":     "Gyppee",
	"You":       "Gyou",
	"I":         "Gi",
}

// Dictionary for words starting with "xr"
var xrDict = map[string]string{
	"xray": "gexray",

	"Xray": "Gexray",
}

// Dictionary for words starting with consonant sounds
var consonantDict = map[string]string{
	"chair":  "airchogo",
	"plynth": "ynthplogo",
	"sea":    "easogo",
	"pogo":   "ogopogo",

	"Chair":  "Airchogo",
	"Plynth": "Ynthplogo",
	"Sea":    "Easogo",
	"Pogo":   "Ogopogo",
}

// Dictionary for words starting with consonant sounds and a subsequent "qu"
var quDict = map[string]string{
	"square": "aresquogo",

	"Square": "Aresquogo",
}

// Test translating words that start with a vowel.
func TestVowels(t *testing.T) {
	for eng, gop := range vowelDict {
		assert.Equal(t, gop, TranslateWord(eng))
	}
}

// Test translating words that start with "xg".
func TestXr(t *testing.T) {
	for eng, gop := range xrDict {
		assert.Equal(t, gop, TranslateWord(eng))
	}
}

// Test translating words that start with a consonant.
func TestConsonants(t *testing.T) {
	for eng, gop := range consonantDict {
		assert.Equal(t, gop, TranslateWord(eng))
	}

	// Test the ones starting with consonant sounds and a "qu" following it
	for eng, gop := range quDict {
		assert.Equal(t, gop, TranslateWord(eng))
	}
}

// Test punctuation.
func TestPunctuation(t *testing.T) {
	assert.Equal(t, "(oggodogo),", TranslateWord("(doggo),"))
}

// Test invalid strings.
func TestInvalid(t *testing.T) {
	assert.Equal(t, "", TranslateWord(""))
	assert.Equal(t, "(gunintelligible)", TranslateWord("don't"))
	assert.Equal(t, "(gunintelligible)", TranslateWord("Wouldn't"))
}

// Test translation of sentences.
func TestSentences(t *testing.T) {
	// Test a simple sentence.
	translation, err := TranslateSentence("I ate the popcorn.")
	assert.Nil(t, err)
	assert.Equal(t, "Gi gate ethogo opcornpogo.", translation)

	// Test a sentence that doesn't end with a supported punctuation symbol.
	_, err = TranslateSentence("I ate the popcorn")
	assert.NotNil(t, err)

	// Test a sentence with a lot of punctuation.
	translation, err = TranslateSentence("I (ate) the, :)popcorn.")
	assert.Nil(t, err)
	assert.Equal(t, "Gi (gate) ethogo, :)opcornpogo.", translation)
}
