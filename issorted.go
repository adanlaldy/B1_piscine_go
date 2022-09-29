package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	var exit bool
	for i := 1; i < len(a); i++ {
		if f(a[i], a[i+1]) == -1 && f(a[i-1], a[i]) == -1 {
			exit = true
		}
		if f(a[i], a[i+1]) == 1 && f(a[i-1], a[i]) == 1 {
			exit = true
		} else {
			break
		}
	}
	return exit
}
