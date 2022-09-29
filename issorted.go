package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	var exit bool
	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) >= 0 {
			exit = true
		} else {
			exit = false
			break
		}
	}
	if exit == false {
		for j := 1; j < len(a); j++ {
			if f(a[j-1], a[j]) <= 0 {
				exit = true
			} else {
				exit = false
				break
			}
		}
	}
	if len(a) == 1 {
		return true
	}
	return exit
}
