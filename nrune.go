package piscine

func NRune(s string, n int) rune {
	i := []rune(s)
	if n > len(s) || n <= 0 {
		return 0
	} else {
		return rune(i[n-1])
	}
}
