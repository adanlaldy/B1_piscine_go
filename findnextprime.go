package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) {
		return nb
	} else if IsPrime(nb) == false {
		for i := nb; IsPrime(nb) == false; i++ {
			nb++
		}
	}
	return nb
}
