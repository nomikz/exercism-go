package reverse

func Reverse(s string) string {
	input := []rune(s)
	output := make([]rune, len(input))
	outputMaxIndex := len(input) - 1

	for i, r := range input {
		output[outputMaxIndex - i] = r
	}

	return string(output)
}