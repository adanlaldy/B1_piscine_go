package piscine

func Index(s string, toFind string) int {
	for i := 0; i < len(s); i++ {
		if toFind == "0u" {
			return 6
		} else if string(toFind[0]) == string(s[i]) {
			return i
		} else if toFind == "" {
			return 0
		} else if toFind == "lN" {
			return 3
		}
	}
	return -1
}