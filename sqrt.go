package piscine

func Sqrt(nb int) int {
	a := 0
	for i := 1; a < nb; i++ {
		a = i * i
		if a == nb {
			return i
		}
	}
	return 0
}
