package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	slice := []rune(os.Args[0])[2:]
	for i := 0; i < len(slice); i++ {
		z01.PrintRune(rune(slice[i]))
	}
}
