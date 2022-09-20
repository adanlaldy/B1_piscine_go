package piscine

func TrimAtoi(s string) int {
	a := 0
	if s == "" {
		return 0
	}
	for i := 0; i < len(s); i++ {
		if int(s[i]) >= 49 && int(s[i]) <= 57 {
			a = a*10 + (int(s[i]) - 48)
		}
	}
	for i := 1; i < len(s); i++ {
		if int(s[i-1]) >= 'a' && int(s[i-1]) <= 'z' && int(s[i]) == '-' {
			a = -a
		}
	}
	return a
}
