package piscine

func Map(f func(int) bool, a []int) []bool {
	slice := []bool{}
	for i := 0; i < len(a); i++ {
		slice = append(slice, f(a[i]))
	}
	return slice
}
