package util

func Runes(row string) []rune {
	runes := []rune(row)

	var result []rune
	for _, r := range runes {
		if r != 65038 {
			result = append(result, r)
		}
	}
	return result
}
