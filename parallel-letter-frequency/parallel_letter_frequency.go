package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(strs []string) FreqMap {
	m := make(FreqMap)
	ch := make(chan FreqMap)

	for _, s := range strs {
		go func(str string) {
			ch <- Frequency(str)
		}(s)
	}

	for range strs {
		fm := <-ch
		for k, v := range fm {
			m[k] += v
		}
	}

	return m
}
