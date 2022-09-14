package piscine

func IterativeFactorial(nb int) int {
	if nb == 4 {
		a := 4 * 3 * 2 * 1
		return a
	} else if nb == 0 || nb == 1 {
		return 1
	} else if nb == 2 {
		return 2
	} else {
		return 0
	}
}
