package anagrammer

import (
	"reflect"
	"testing"
)

func Test_Find(t *testing.T) {
	tests := []struct {
		name      string
		inputWord string
		expected  []string
	}{
		{
			name:      "with no input",
			inputWord: "",
			expected:  []string{},
		},
		{
			name:      "with single letter input",
			inputWord: "a",
			expected:  []string{},
		},
		{
			name:      "with no anagrams",
			inputWord: "racecar",
			expected:  []string{},
		},
		{
			name:      "with anagrams",
			inputWord: "battle",
			expected:  []string{"battel", "tablet"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Find(tt.inputWord)

			if reflect.DeepEqual(tt.expected, actual) == false {
				t.Errorf("\nexpected: %q\n     got: %q", tt.expected, actual)
			}
		})
	}
}
