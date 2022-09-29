package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	var exit bool
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) == 1 {
			exit = true
		}
		for j := 0; j < len(a)-1; j++ {
			if f(a[j], a[j+1]) == -1 {
				exit = true
			}
		}
	}
	return exit
}
