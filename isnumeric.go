package piscine

func IsNumeric(s string) bool {
	for i := 0; i < len(s); i++ {
		if int(s[i]) >= 48 && int(s[i]) <= 57 || int(s[i]) >= 97 && int(s[i]) <= 122 {
			continue
		} else {
			return false
		}
	}
	return true
}
