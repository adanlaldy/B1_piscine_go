package piscine

func AppendRange(min, max int) []int {
	slice := []int{min}
	for i := min; i < max-1; i++ {
		slice = append(slice, i+1)
	}
	if min < max {
		return slice
	} else if min >= max {
		return nil
	}
	return slice
}
