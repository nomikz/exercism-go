package raindrops

import (
	"sort"
	"strconv"
)

var numToStr = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

func Convert(inputNum int) string {
	var resString string
	notAny := true

	keys := make([]int, 0)
	for k := range numToStr {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, v := range keys {
		if inputNum % v == 0 {
			notAny = false

			resString += numToStr[v]
		}
	}

	if notAny {
		resString = strconv.Itoa(inputNum)
	}

	return resString
}