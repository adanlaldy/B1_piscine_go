package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for k := 1; k < len(os.Args); k++ {
		slice := []rune(os.Args[k])
		for i := 0; i < len(slice); i++ {
			z01.PrintRune(rune(slice[i]))
		}
		z01.PrintRune('\n')
	}
}
