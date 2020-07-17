package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const (
	dictionaryPath = "/usr/share/dict/words"
)

var inputWord string

func init() {
	flag.StringVar(&inputWord, "word", "", "the word to anagram")
	flag.StringVar(&inputWord, "w", "", "the word to anagram (short hand)")
}

func main() {
	// Accept input on the command line
	flag.Parse()

	// Create the word list from the internal dictionary
	wordList, err := buildWordList(dictionaryPath)
	if err != nil {
		log.Fatal(err)
	}

	anagrams := findAnagrams(inputWord, wordList)

	if len(anagrams) == 0 {
		fmt.Printf("'%s' has no anagrams\n", inputWord)
	} else {
		outputWords := strings.Join(anagrams, ", ")

		fmt.Printf("Found %d anagrams for '%s':\n", len(anagrams), inputWord)
		fmt.Println(outputWords)
	}

	os.Exit(0)
}

/* -------------------- Unexported Functions -------------------- */

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

// buildWordList uses the local dictionary to create an enumerable set of all known
// words. This is used to check against for anagrams
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
