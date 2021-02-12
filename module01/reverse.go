package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func ReverseLame(word string) string {
	size := len(word)
	if size <= 1 {
		return word
	}
	return word[size-1:] + Reverse(word[:size-1])
}

func Reverse(word string) string {
	var result string
	for _, rune := range word {
		result = string(rune) + result
	}
	return result
}
