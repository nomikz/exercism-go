package scrabble

import (
	"strings"
)

var charScore = make(map[rune]int)

func Score(s string) int {
	rawScore := map[string]int{
		"aeioulnrst": 1,
		"dg":         2,
		"bcmp":       3,
		"fhvwy":      4,
		"k":          5,
		"jx":         8,
		"qz":         10,
	}

	for letters, score := range rawScore {
		for _, letter := range letters {
			charScore[letter] = score
		}
	}

	var score int

	for _, c := range strings.ToLower(s) {
		score += charScore[c]
	}

	return score
}
