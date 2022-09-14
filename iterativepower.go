package piscine

func IterativePower(nb int, power int) int {
	if nb < 0 {
		return 0
	} else {
		nb = nb * IterativePower(nb, nb*nb*nb)
	}
	return nb
}
