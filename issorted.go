package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	var exit bool
	for i := 1; i < len(a)-6; i++ {
		if f(a[i], a[i+1]) == -1 && f(a[i-1], a[i]) == -1 && f(a[i+1], a[i+2]) == -1 && f(a[i+2], a[i+3]) == -1 && f(a[i+3], a[i+4]) == -1 && f(a[i+4], a[i+5]) == -1 && f(a[i+5], a[i+6]) == -1 {
			exit = true
		} else if f(a[i], a[i+1]) == 1 && f(a[i-1], a[i]) == 1 && f(a[i+1], a[i+2]) == 1 && f(a[i+2], a[i+3]) == 1 && f(a[i+3], a[i+4]) == 1 && f(a[i+4], a[i+5]) == 1 && f(a[i+5], a[i+6]) == 1 {
			exit = true
		}
	}
	return exit
}
