package piscine

func StrRev(s string) string {
	array := []rune(s)
	var a string
	for i := len(array) - 1; i >= 0; i-- {
		a += string(array[i])
	}
	return a
}
