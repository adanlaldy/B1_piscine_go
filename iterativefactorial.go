package piscine

func IterativeFactorial(nb int) int {
	if nb == 1 || nb == 0 {
		return 1
	} else if nb > 65535 {
		return 0
	} else {
		nb = nb * (IterativeFactorial(nb - 1))
	}
	return nb
}
