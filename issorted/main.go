package main

import (
	"fmt"
	"piscine"
)

func f(a, b int) int {
	vr := 0
	if a > b {
		vr = 1
	} else if a == b {
		vr = 0
	} else if a < b {
		vr = -1
	}
	return vr
}

func main() {
	a1 := []int{922963, 160821, -355951, -488076, 903098, 601504, -265473, -989152}
	a2 := []int{-952730, -920742, -613126, -446098, 16800, 331325, 415212, 571045}

	result1 := piscine.IsSorted(f, a1)
	result2 := piscine.IsSorted(f, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
