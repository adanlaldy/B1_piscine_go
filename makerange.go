package piscine

func MakeRange(min, max int) []int {
	sliceempty := []int{0}
	if min >= max {
		return sliceempty
	}
	slice := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		slice[i] = min + i
	}
	return slice
}
