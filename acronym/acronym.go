package acronym

import (
	"strings"
	"unicode/utf8"
)

func Abbreviate(s string) string {
	var resString string

	s = strings.Replace(s, "-", " ", -1)

	words := strings.Split(s, " ")

	for _, w := range words {
		if utf8.RuneCountInString(w) == 0 {
			continue
		}

		for _, char := range []string{"_", ","} {
			w = strings.Trim(w, char)
		}

		resString += w[:1]
	}

	return strings.ToUpper(resString)
}
