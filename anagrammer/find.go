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
	anagrams, err := findAnagrams(dictionaryPath, inputWord)
	if err != nil {
		log.Fatal(err)
	}

	return anagrams
}

/* -------------------- Unexported Functions -------------------- */

// findAnagrams uses the local dictionary to create an enumerable set of all known
// words. This is used to check for anagrams
func findAnagrams(dictPath string, inputWord string) ([]string, error) {
	words := []string{}

	dict, err := os.Open(dictPath)
	if err != nil {
		return words, err
	}

	scanner := bufio.NewScanner(dict)
	for scanner.Scan() {
		word := scanner.Text()

		if len(word) != len(inputWord) {
			// If the word we're scanning is a different length than the input word,
			// this word cannot be an anagram of the input word, so don't include it in the dict
			continue
		}

		if word == inputWord {
			// If the word is the same word as the input word, it is not an anagram
			continue
		}

		isAnagram := compareForEquality(inputWord, word)
		if isAnagram {
			words = append(words, word)
		}
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
