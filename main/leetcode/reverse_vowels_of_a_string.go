package leetcode

func ReverseVowels(s string) string {
	chars := []rune(s)
	i, j := 0, len(chars)-1

	for i < j {
		if isVowel(chars[i]) && isVowel(chars[j]) {
			chars[i], chars[j] = chars[j], chars[i]
			i++
			j--
		} else if !isVowel(chars[i]) {
			i++
		} else {
			j--
		}
	}

	return string(chars)
}

func isVowel(c rune) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
		c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}
