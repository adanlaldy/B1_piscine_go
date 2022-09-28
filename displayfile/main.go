package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Print("File name missing")
		return
	} else if len(os.Args) > 2 {
		fmt.Print("Too many arguments")
	} else {
		a, _ := os.ReadFile("quest8.txt")
		fmt.Print(string(a))
	}
}
