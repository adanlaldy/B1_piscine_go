package piscine

func BasicAtoi2(s string) int {
	a := 0
	for i := 0; i < len(s); i++ {
		a = a*10 + (int(s[i]) - 48)
		if int(s[i]) < 48 || int(s[i]) > 57 {
			return 0
		}
	}
	return a
}
