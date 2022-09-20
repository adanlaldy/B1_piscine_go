package piscine

func ToLower(s string) string {
	slice := []rune(s)
	for i := 0; i < len(s); i++ {
		if rune(s[i]) >= 65 && rune(s[i]) <= 90 {
			slice[i] = rune(s[i]) + 32
		} else if rune(s[i]) >= 97 && rune(s[i]) <= 122 {
			continue
		}
	}
	return string(slice)
}
