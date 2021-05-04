package hamming

import "errors"

func Distance(a, b string) (int, error) {
	var distance int

	if len(a) != len(b) {
		return distance, errors.New("lengths' of strings are not equal")
	}


	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
