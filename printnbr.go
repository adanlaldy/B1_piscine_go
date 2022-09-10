package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	var nb int = n
	if nb > 10 {
		PrintNbr(nb / 10)
	} else if nb < 0 {
		z01.PrintRune(rune(45))
		nb = -nb
		PrintNbr(nb / 10)
	}
	z01.PrintRune(rune((nb % 10) + 48))
}
