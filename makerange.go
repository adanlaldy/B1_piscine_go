package piscine

func MakeRange(min, max int) []int {
	slice := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		slice[i] = min + i
	}
	if min < max {
		return slice
	} else if min >= max {
		return nil
	}
	return slice
}
