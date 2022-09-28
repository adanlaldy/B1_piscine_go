package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("File name missing")
		return
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	} else {
		a, _ := os.ReadFile("quest8.txt")
		fmt.Println(string(a))
	}
}
