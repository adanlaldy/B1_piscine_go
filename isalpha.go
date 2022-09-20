package piscine

func IsAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		if int(s[i]) >= 97 && int(s[i]) <= 122 || int(s[i]) >= 65 && int(s[i]) <= 90 || s == "" || int(s[i]) >= 48 && int(s[i]) <= 57 {
			continue
		} else {
			return false
		}
	}
	return true
}
