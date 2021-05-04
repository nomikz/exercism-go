package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

var onlyNums = regexp.MustCompile(`^[0-9]+$`)

func Valid(s string) bool {
	s = strings.Join(strings.Fields(s), "")

	if !onlyNums.Match([]byte(s)) {
		return false
	}

	if len(s) <= 1 {
		return false
	}

	var bs []byte
	for _, v := range strings.Split(s, "") {
		integerV, _ := strconv.Atoi(v)

		bs = append(bs, byte(integerV))
	}

	newS := make([]byte, len(bs))
	isRecalc := false
	for i := len(bs) - 1; i >= 0; i-- {
		if isRecalc {
			p := int(bs[i]) * 2

			if p > 9 {
				p -= 9
			}

			newS[i] = byte(p)
		} else {
			newS[i] = bs[i]
		}
		isRecalc = !isRecalc
	}

	var sum int
	for _, b := range newS {
		sum += int(b)
	}

	return sum % 10 == 0
}
