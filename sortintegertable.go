package piscine

import "fmt"

func SortIntegerTable(table []int) int {
	a := 0
	for i := 0; i < len(table); i++ {
		for j := i + 1; j < len(table); j++ {
			if table[j] < table[a] {
				a = j
			}
		}
	}
	fmt.Println(table[a])
	return (table[a])
}
