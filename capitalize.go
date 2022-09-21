package piscine

func Capitalize(s string) string {
	slice := []rune(s)
	for i := 0; i < len(s); i++ {
		if rune(s[i]) >= 97 && rune(s[i]) <= 122 {
			continue
		} else if rune(s[i]) >= 97 && rune(s[i]) <= 122 && (rune(s[i-1]) >= 65 && rune(s[i-1]) <= 90 || rune(s[i-1]) == 32 || rune(s[i-1]) == 43) {
			slice[i] = rune(s[i]) - 32
		}
	}
	return string(slice)
}
