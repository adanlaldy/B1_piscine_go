package piscine

func IsUpper(s string) bool {
	a := 0
	for i := 0; i < len(s); i++ {
		a = i
	}
	if int(s[a]) >= 65 && int(s[a]) <= 90 {
		return true
	} else {
		return false
	}
	for j := 0; j < len(s); j++ {
		if int(s[j]) >= 65 && int(s[j]) <= 90 {
			return true
		} else {
			return false
		}
		return false
	}
	return false
}
