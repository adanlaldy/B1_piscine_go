package piscine

func IsPrintable(s string) bool {
	for i := 0; i < len(s); i++ {
		if int(s[i]) >= 48 && int(s[i]) <= 57 || int(s[i]) >= 97 && int(s[i]) <= 122 || int(s[i]) >= 65 && int(s[i]) <= 90 {
			continue
		} else {
			return false
		}
	}
	return true
}
