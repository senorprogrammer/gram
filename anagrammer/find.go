package anagrammer

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

const (
	dictionaryPath = "/usr/share/dict/words"
)

// Find takes an input word and returns a set of all the anagrams for that word
func Find(inputWord string) []string {
	wordList, err := buildWordList(dictionaryPath)
	if err != nil {
		log.Fatal(err)
	}

	anagrams := findAnagrams(inputWord, wordList)

	return anagrams
}

/* -------------------- Unexported Functions -------------------- */

// buildWordList uses the local dictionary to create an enumerable set of all known
// words. This is used to check for anagrams
func buildWordList(dictPath string) ([]string, error) {
	words := []string{}

	dict, err := os.Open(dictPath)
	if err != nil {
		return words, err
	}

	scanner := bufio.NewScanner(dict)
	for scanner.Scan() {
		word := scanner.Text()
		words = append(words, word)
	}

	return words, nil
}

// compareForEquality checks whether or not two strings, when sorted, are exactly the same
func compareForEquality(a, b string) bool {
	if len(a) != len(b) {
		// If they don't have the same length, they cannot be the same
		return false
	}

	aSliced := strings.Split(a, "")
	bSliced := strings.Split(b, "")

	sort.Sort(StringSlice(aSliced))
	sort.Sort(StringSlice(bSliced))

	for idx, aLetter := range aSliced {
		bLetter := bSliced[idx]

		if aLetter != bLetter {
			// If any letter differs positionally, they cannot be the same
			return false
		}
	}

	return true
}

// findAnagrams finds anagrams for a given string
func findAnagrams(inputWord string, wordList []string) []string {
	anagrams := []string{}

	if len(inputWord) <= 1 {
		// A one-letter string, or a no-letter string, cannot have any anagrams
		return anagrams
	}

	// Check every word in the word list to see if it's an anagram
	// The following conditions exclude a word from being an anagram:
	//	* if the word is a different length
	//  * if the word is the same word as the input word
	//  * if the word contains a letter that the input word does not
	for _, dictWord := range wordList {
		// Check to see if the word is a different length
		dictWordLen := len(dictWord)
		if dictWordLen != len(inputWord) {
			continue
		}

		// The word itself is not an anagram of itself
		if dictWord == inputWord {
			continue
		}

		// Compare the two words to see if they have any differing letters
		// If any letters are different, they cannot be anagrams
		areTheSame := compareForEquality(inputWord, dictWord)

		if areTheSame {
			anagrams = append(anagrams, dictWord)
		}
	}

	return anagrams
}
