package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args
	for k := len(a) - 1; k > 0; k-- {
		slice := []rune(os.Args[k])
		for i := 0; i < len(slice); i++ {
			z01.PrintRune(rune(slice[i]))
		}
		z01.PrintRune('\n')
	}
}
