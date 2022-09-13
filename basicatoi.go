package piscine

func BasicAtoi(s string) int {
	a := 0
	for i := 0; i < len(s); i++ {
		a = a*10 + (int(s[i]) - 48)
	}
	return a
}
