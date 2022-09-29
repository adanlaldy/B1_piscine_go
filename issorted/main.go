package main

import (
	"fmt"
	"piscine"
)

func f(a, b int) int {
	exit := 0
	if a > b {
		exit = 1
	} else if a == b {
		exit = 0
	} else {
		exit = -1
	}
	return exit
}

func main() {
	a1 := []int{694269, 638844, 579652, 409813, -144530, -410537, -639118, -972120}
	a2 := []int{0, 2, 1, 3}

	result1 := piscine.IsSorted(f, a1)
	result2 := piscine.IsSorted(f, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
