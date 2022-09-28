package piscine

func CountIf(f func(string) bool, tab []string) int {
	var b int
	for i := 0; i < len(tab); i++ {
		f(tab[i])
		if f(tab[i]) == true {
			b++
		}
	}
	return b
}
