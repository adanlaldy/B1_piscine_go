package piscine

func IsPrintable(s string) bool {
	for i := 0; i < len(s); i++ {
		if int(s[i]) >= 32 && int(s[i]) <= 126 || int(s[i]) >= 128 && int(s[i]) <= 254 {
			continue
		} else {
			return false
		}
	}
	return true
}
