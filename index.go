package piscine

func Index(s string, toFind string) int {
	for i := 0; i < len(s); i++ {
		if toFind[0] == s[i] {
			return i
		}
	}
	return -1
}
