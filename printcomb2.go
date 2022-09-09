package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for a := 0; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			for c := 0; c <= 9; c++ {
				for d := 0; d <= 9; d++ {
					if a*100+b < c*100+d {
						z01.PrintRune(rune(48 + a))
						z01.PrintRune(rune(48 + b))
						z01.PrintRune(' ')
						z01.PrintRune(rune(48 + c))
						z01.PrintRune(rune(48 + d))
						if a == 9 && b == 8 && c == 9 && d == 9 {
							z01.PrintRune('\n')
						} else {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
}
