package piscine

func AlphaCount(s string) int {
	a := 0
	for i := 0; i < len(s); i++ {
		if rune(s[i]) >= 97 && rune(s[i]) <= 122 || rune(s[i]) >= 65 && rune(s[i]) <= 90 {
			a++
		}
	}
	return a
}
