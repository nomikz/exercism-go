package isogram

import "strings"

func IsIsogram(s string) bool {
	m := make(map[rune]bool)

	if len(s) == 0 {
		return true
	}

	for _, r := range strings.ToLower(s) {
		_, ok := m[r]

		if r == ' ' || r == '-' {
			continue
		}

		if ok {
			return false
		}

		m[r] = true
	}

	return true
}