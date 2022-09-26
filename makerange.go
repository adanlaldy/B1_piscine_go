package piscine

func MakeRange(min, max int) []int {
	slice := make([]int, min)
	for i := 0; i < len(slice); i++ {
		slice[i] = min + i
	}
	if min < max {
		return slice
	} else if min >= max {
		return nil
	}
	return slice
}
