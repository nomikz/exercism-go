package bob

import (
	"regexp"
	"strings"
	"unicode"
)

func Hey(remark string) string {
	if isAllCaps(remark) && strings.HasSuffix(strings.TrimSpace(remark), "?") {
		return "Calm down, I know what I'm doing!"
	}

	if strings.HasSuffix(strings.TrimSpace(remark), "?") {
		return "Sure."
	}

	if isAllCaps(remark) {
		return "Whoa, chill out!"
	}

	if len(remark) == 0 || isAllSpace(remark) {
		return "Fine. Be that way!"
	}

	return "Whatever."
}

func isAllSpace(s string) bool {
	for _, v := range s {
		if !unicode.IsSpace(v) {
			return false
		}
	}

	return true
}

func isAllCaps(s string) bool {
	hasUnicodeStr := false

	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString

	for _, v := range strings.Split(s, "") {
		if isAlpha(v) {
			hasUnicodeStr = true
		}

		if v != strings.ToUpper(v) {
			return false
		}
	}

	return hasUnicodeStr
}