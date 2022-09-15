package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else if power > 1 {
		nb = nb * RecursivePower(nb, power-1)
	} else if nb == 64 {
		return nb
	}
	return nb
}
