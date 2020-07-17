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
	anagrams, err := scanForAnagrams(dictionaryPath, inputWord)
	if err != nil {
		log.Fatal(err)
	}

	return anagrams
}

/* -------------------- Unexported Functions -------------------- */

// scanForAnagrams uses the local dictionary to create an enumerable set of all known
// words. This is used to check for anagrams
func scanForAnagrams(dictPath string, inputWord string) ([]string, error) {
	anagrams := []string{}

	dict, err := os.Open(dictPath)
	if err != nil {
		return anagrams, err
	}

	scanner := bufio.NewScanner(dict)
	for scanner.Scan() {
		word := scanner.Text()

		if wordsAreEqual(inputWord, word) {
			anagrams = append(anagrams, word)
		}
	}

	return anagrams, nil
}

// wordsAreAnagrams checks whether or not two strings, when sorted, are exactly the same
func wordsAreEqual(wordA, wordB string) bool {
	if len(wordA) != len(wordB) {
		// If the word we're scanning is a different length than the input word,
		// this word cannot be anagrams
		return false
	}

	if wordA == wordB {
		// If the word is the same word as the input word, it is not an anagram
		return false
	}

	aSliced := strings.Split(wordA, "")
	bSliced := strings.Split(wordB, "")

	// Convert the slices to StringSlice so they can be sorted
	sort.Sort(StringSlice(aSliced))
	sort.Sort(StringSlice(bSliced))

	for idx, aLetter := range aSliced {
		bLetter := bSliced[idx]

		if aLetter != bLetter {
			// If any letter differs positionally, they cannot be anagrams
			return false
		}
	}

	return true
}
