package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	c := true
	for i := 0; i < len(a)-1; i++ {
		if int(a[i]) < int(a[i+1]) {
			continue
		} else {
			break
		}
	}
	for j := 0; j < len(a)-1; j++ {
		if int(a[j]) > int(a[j+1]) {
			continue
		} else {
			c = false
			break
		}
	}
	return c
}
