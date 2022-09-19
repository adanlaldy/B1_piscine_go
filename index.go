package piscine

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}
	S := AlphaCount(s)
	I := AlphaCount(toFind)
	for i := 0; I <= S; i++ {
		if Compare(s[i:I], toFind) == 0 {
			return i
		}
		I++
	}
	return -1
}
