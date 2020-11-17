package translator

import (
	"fmt"
	"strings"
)

/*
The definitive grammar guide on English to Gopherish translation:

1. If a word starts with a vowel letter, add prefix “g” to the word (ex. apple => gapple)
2. If a word starts with the consonant letters “xr”, add the prefix “ge” to the begging of the word.
Such words as “xray” actually sound in the beginning with vowel sound as you pronounce them so a true gopher would say “gexray”.
3. If a word starts with a consonant sound, move it to the end of the word and then add “ogo” suffix to the word.
Consonant sounds can be made up of multiple consonants, a.k.a. a consonant cluster (e.g. "chair" -> "airchogo”).
4. If a word starts with a consonant sound followed by "qu", move it to the end of the word, and then add "ogo" suffix to the word (e.g. "square" -> "aresquogo").
*/

// Use map for easier "contains" checks.
// All known vowels.
var vowels = map[string]interface{}{"a": 0, "e": 0, "i": 0, "o": 0, "u": 0, "y": 0}

// All known consonants.
var consonants = map[string]interface{}{"b": 0, "c": 0, "d": 0, "f": 0, "g": 0, "h": 0, "j": 0, "k": 0, "l": 0, "m": 0,
	"n": 0, "p": 0, "q": 0, "r": 0, "s": 0, "t": 0, "v": 0, "w": 0, "x": 0, "z": 0}

// Capitals of all known letters.
var capitals = map[string]interface{}{"B": 0, "C": 0, "D": 0, "F": 0, "G": 0, "H": 0, "J": 0, "K": 0, "L": 0, "M": 0,
	"N": 0, "P": 0, "Q": 0, "R": 0, "S": 0, "T": 0, "V": 0, "W": 0, "X": 0, "Z": 0, "Y": 0, "A": 0, "E": 0, "I": 0,
	"O": 0, "U": 0}

// TranslateWord translates a word from English to Gopherish.
func TranslateWord(word string) string {
	// Empty in English should be empty in Gopherish.
	if word == "" {
		return ""
	}

	// Skip translating words with apostrophes.
	if strings.Contains(word, "'") {
		return "(gunintelligible)"
	}

	// Strip word from punctuation.
	leading, word, trailing := stripPunctuation(word)

	// Determine if word is capitalized.
	_, isCapital := capitals[word[:1]]

	// Work with lowercase for consistency.
	word = strings.ToLower(word)
	first := word[:1]

	// Output word in Gopherish.
	var translated = ""

	// Handle words starting with a vowel.
	if _, ok := vowels[first]; ok {
		translated = prefixG(word)
	}

	// Handle words starting with "xr"...
	if len(word) >= 2 && word[0:2] == "xr" {
		translated = prefixGe(word)
		// ...and handle all other consonant sounds.
	} else if _, ok := consonants[first]; ok {
		translated = postfixOgo(extractConsonantSound(word))
	}

	// Capitalize if necessary.
	// Assume capitalized words in English are capitalized in Gopherish, too.
	if isCapital {
		translated = strings.Title(translated)
	}

	translated = fmt.Sprintf("%s%s%s", leading, translated, trailing)

	// Make the user aware of words that cannot be translated, such as ones starting with unhandled symbols.
	if translated == "" {
		translated = "(gunintelligible)"
	}
	return translated
}

// extractConsonantSound returns a cluster of consonant sounds and the remaining base of the word.
// Assume consonant sound is any sequence of consonants
// (diverging from actual definition as per requirements and for simplicity)
func extractConsonantSound(word string) (string, string) {
	var conSound = ""

	// Append consonants to conSound.
	for {
		if word == "" {
			break
		}
		first := word[0:1]
		_, hasCon := consonants[first]
		if !hasCon {
			break
		}
		conSound = fmt.Sprintf("%s%s", conSound, first)
		word = word[1:]
	}

	// Append "u" if last consonant was "q".
	if conSound[len(conSound)-1:] == "q" && word[:1] == "u" {
		conSound = fmt.Sprintf("%su", conSound)
		word = word[1:]
	}

	return conSound, word
}

// prefixG prefixes "g" to a word.
// Assume word begins with a vowel.
func prefixG(word string) string {
	return fmt.Sprintf("g%s", word)
}

// prefixGe prefixes "ge" to a word.
// Assume word begins with 'xr'
func prefixGe(word string) string {
	return fmt.Sprintf("ge%s", word)
}

// postfixOgo postfixes "ogo" to a word.
// Assume word starts with a consonant sound (1 or more consonant as per specification) and eventually following "qu".
func postfixOgo(consonants string, base string) string {
	return fmt.Sprintf("%s%sogo", base, consonants)
}

// TranslateSentence translates a sentence from English to Gopherish.
func TranslateSentence(sentence string) (string, error) {
	ending := sentence[len(sentence)-1:]
	if ending != "!" && ending != "." && ending != "?" {
		return "", fmt.Errorf("invalid sentence ending in '%s'.\nOnly '.', '?' and '!' are supported", sentence)
	}
	words := strings.Fields(sentence)

	// Strip word of leading and trailing punctuation, translate it and then reassign with original punctuation.
	for i, word := range words {
		words[i] = TranslateWord(word)
	}

	return strings.Join(words, " "), nil
}

// stripPunctuation separates leading and trailing punctuation from words.
// First return value is a set of leading punctuation.
// Second return value is the extracted word.
// Third return value is a set of trailing punctuation.
func stripPunctuation(word string) (string, string, string) {
	var leading = ""
	var trailing = ""

	// Strip leading punctuation.
	for range word {
		if word == "" {
			break
		}
		first := word[:1]
		if !isALetter(first) {
			word = word[1:]
			leading = fmt.Sprintf("%s%s", leading, first)
		} else {
			break
		}
	}

	// Strip trailing punctuation
	for range word {
		if word == "" {
			break
		}
		last := word[len(word)-1:]
		if !isALetter(last) {
			word = word[:len(word)-1]
			trailing = fmt.Sprintf("%s%s", last, trailing)
		} else {
			break
		}
	}

	return leading, word, trailing
}

// isALetter sifts out letters from punctuation.
// Assume all non-letters are punctuation.
func isALetter(symbol string) bool {
	if _, isVowel := vowels[symbol]; isVowel {
		return true
	}
	if _, isConsonant := consonants[symbol]; isConsonant {
		return true
	}
	if _, isCapital := capitals[symbol]; isCapital {
		return true
	}
	return false
}
