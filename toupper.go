package piscine

func ToUpper(s string) string {
	slice := []rune(s)
	for i := 0; i < len(s); i++ {
		if rune(s[i]) >= 97 && rune(s[i]) <= 122 {
			slice[i] = rune(s[i]) - 32
		} else if rune(s[i]) >= 65 && rune(s[i]) <= 90 {
			continue
		}
	}
	return string(slice)
}
