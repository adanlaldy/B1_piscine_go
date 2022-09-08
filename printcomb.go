package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for a := 0; a != 10; a++ {
		for b := 1; b != 10; b++ {
			for c := 2; c != 10; c++ {
				if a < b && b < c {
					z01.PrintRune(rune(48 + a))
					z01.PrintRune(rune(48 + b))
					z01.PrintRune(rune(48 + c))
					if a != 7 {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
