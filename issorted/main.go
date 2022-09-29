package main

import (
	"fmt"
	"piscine"
)

func f(a, b int) int {
	return a - b
	//vr := 0
	//if a > b {
	//	vr = 1
	//} else if a == b {
	//	vr = 0
	//} else if a < b {
	//	vr = -1
	//}
	//return vr
}

func main() {
	a1 := []int{0}
	a2 := []int{1, 2, 3, 4, 5, 6}

	result1 := piscine.IsSorted(f, a1)
	result2 := piscine.IsSorted(f, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
