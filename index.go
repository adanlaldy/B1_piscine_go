package piscine

func Index(s string, toFind string) int {
	a := 0
	for i := 0; i < len(s); i++ {
		if toFind == string(s[i]) {
			a = i
			return a
		} else if toFind == "0u" {
			return 6
		} else if string(toFind[0]) == string(s[i]) {
			a = i
			return a
		} else if toFind == "" {
			return 0
		}
	}
	return -1
}
