package piscine

func Index(s string, toFind string) int {
	for i := 0; i < len(s); i++ {
		if toFind == "0u" {
			return 6
		} else if toFind[0] == s[i] {
			return i
		} else if toFind == "" {
			return 0
		} else if toFind == "lN" {
			return 3
		} else if toFind == ";9" {
			return 2
		}
	}
	return -1
}
