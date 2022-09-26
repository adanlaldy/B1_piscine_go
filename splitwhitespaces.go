package piscine

func SplitWhiteSpaces(s string) []string {
	slice := []string{}
	char := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\t' || s[i] == '\n' {
			slice = append(slice, char)
			char = ""
		} else {
			char += string(s[i])
		}
	}
	slice = append(slice, char)
	return slice
}
