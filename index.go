package piscine

func Index(s string, toFind string) int {
	a := 0
	for i := 0; i < len(s); i++ {
		if toFind == string(s[i]) {
			return i
			a = i
		} else if string(toFind[0]) == string(s[i]) {
			return i
			a = i
		} else if string(toFind[0]) < string(s[i]) {
			return -1
		}
	}
	return a
}
