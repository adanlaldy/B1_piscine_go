package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	var c bool
	for i := 0; i < len(a)-1; i++ {
		if int(a[i]) < int(a[i+1]) {
			c = true
		} else {
			c = false
			break
		}
	}
	return c
}
