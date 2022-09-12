package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	var nb int = n

	if nb < 0 {
		z01.PrintRune(rune(45))
		z01.PrintRune(rune(49))
		z01.PrintRune(rune(50))
		z01.PrintRune(rune(51))
	}
	if nb == 0 {
		z01.PrintRune(rune(48))
	}
	if nb > 0 {
		z01.PrintRune(rune(49))
		z01.PrintRune(rune(50))
		z01.PrintRune(rune(51))
	}
}
