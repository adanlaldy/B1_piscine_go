package piscine

func IsUpper(s string) bool {
	for i := 0; i < len(s); i++ {
		if int(s[i]) >= 65 && int(s[i]) <= 90 {
			continue
		} else {
			return false
		}
	}
	return true
}
