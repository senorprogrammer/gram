package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/senorprogrammer/gram/anagrammer"
)

const (
	dictionaryPath = "/usr/share/dict/words"
)

// inputWord captures the word from stdin that we'll be finding anagrams for
var inputWord string

func init() {
	flag.StringVar(&inputWord, "word", "", "the word to anagram")
	flag.StringVar(&inputWord, "w", "", "the word to anagram (short hand)")
}

func main() {
	flag.Parse()

	anagrams := anagrammer.Find(inputWord)

	if len(anagrams) == 0 {
		fmt.Printf("'%s' has no anagrams\n", inputWord)
	} else {
		output := strings.Join(anagrams, ", ")

		fmt.Printf("Found %d anagrams for '%s':\n", len(anagrams), inputWord)
		fmt.Println(output)
	}

	os.Exit(0)
}
