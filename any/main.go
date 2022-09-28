package main

import (
	"fmt"
	"piscine"
)

func main() {
	a1 := []string{"vvdxfurouljpi", "pkzgqtamvdrpc", "tweazhmacreri", "suycfotslpvnt", "bkalcjbwiqfff", "gvpktmafylsfq", "vmzqihhekcjjo", "pxobfdjonglyi"}
	a2 := []string{"This", "is", "4", "you"}

	result1 := piscine.Any(piscine.IsNumeric, a1)
	result2 := piscine.Any(piscine.IsNumeric, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
