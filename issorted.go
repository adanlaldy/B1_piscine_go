package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	var exit bool
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) == 1 {
			exit = true
		}
		if f(a[i], a[i+1]) == -1 {
			exit = true
		} else {
			exit = false
		}
	}
	return exit
}
