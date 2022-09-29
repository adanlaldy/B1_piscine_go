package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	var exit bool
	for i := 1; i < len(a)-3; i++ {
		if f(a[i], a[i+1]) == -1 && f(a[i-1], a[i]) == -1 && f(a[i+1], a[i+2]) == -1 && f(a[i+2], a[i+3]) == -1 {
			exit = true
		} else if f(a[i], a[i+1]) == 1 && f(a[i-1], a[i]) == 1 && f(a[i+1], a[i+2]) == 1 && f(a[i+2], a[i+3]) == 1 {
			exit = true
		}
	}
	return exit
}
