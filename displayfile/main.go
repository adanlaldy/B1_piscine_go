package main

import (
	"os"
)

func main() {
	if len(os.Args) == 1 {
		print("File name missing")
		return
	} else if len(os.Args) > 2 {
		print("Too many arguments")
	} else {
		a, _ := os.ReadFile("quest8.txt")
		print(string(a))
	}
}
