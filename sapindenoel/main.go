package main

import (
	"os"
	"piscine"
)

func b(height int) {
	sap := ("*")
	temp := height - 1
	tronk := ("|||")
	for j := 0; j < height; j++ {
		print(string(" "))
	}
	println(sap)
	for i := 0; i < height-1; i++ {
		for j := 0; j < temp; j++ {
			print(string(" "))
		}
		temp--
		sap += ("**")
		println(sap)
	}
	for j := 0; j < height-1; j++ {
		print(string(" "))
	}
	print(tronk)
}
func main() {
	arg := os.Args[1:]
	b(piscine.Atoi(arg[0]))
}
