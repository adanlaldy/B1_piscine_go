package piscine

func Capitalize(s string) string {
	slice := []rune(s)
	for i := 1; i < len(s); i++ {
		if (rune(s[i]) >= 97 && rune(s[i]) <= 122) && !((rune(s[i-1]) >= 97 && rune(s[i-1]) <= 122) || !(rune(s[i+1]) >= 97 && rune(s[i+1]) <= 122)) {
			slice[i] = rune(s[i]) - 32
		}
	}
	for j := 1; j < len(s); j++ {
		if (rune(s[j]) >= 65 && rune(s[j]) <= 90) && !((rune(s[j-1]) >= 97 && rune(s[j-1]) <= 122) || !(rune(s[j+1]) >= 97 && rune(s[j+1]) <= 122)) {
			slice[j] = rune(s[j]) + 32
		}
	}
	return string(slice)
}
