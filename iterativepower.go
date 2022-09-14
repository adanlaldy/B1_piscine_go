package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else if power > 1 {
		nb = nb * IterativePower(nb, power-1)
	} else if nb == 64 {
		return nb
	}
	return nb
}
