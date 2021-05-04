package grains

import (
	"errors"
	"math"
)

func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("invalid value")
	}

	return uint64(math.Pow(2, float64(n - 1))), nil
}

func Total() uint64 {
	var sum uint64
	ch := make(chan uint64)

	for i := 1; i < 65; i++ {
		go func(channel chan uint64, n int) {
			res, _ := Square(n)
			channel <- res
		}(ch, i)
	}

	for i := 1; i < 65; i++ {
		sum += <-ch
	}
	return sum
}
