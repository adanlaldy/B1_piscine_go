package piscine

func Compare(a, b string) int {
	if a == b {
		return 0
	} else if b < a {
		return 1
	} else if b > a {
		return -1
	}
	return 0
}
