package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	slice := []rune(os.Args[1])[0:]
	for i := 0; i < len(slice); i++ {
		z01.PrintRune(rune(slice[i]))
	}
	z01.PrintRune('\n')
}
