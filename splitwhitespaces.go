package piscine

func SplitWhiteSpaces(s string) []string {
	slice := []string{}
	char := ""
	if len(s) == 0 {
		return slice
	}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\t' || s[i] == '\n' {
			if len(char) != 0 {
				slice = append(slice, char)
				char = ""
			}
		} else {
			char += string(s[i])
		}
	}
	slice = append(slice, char)
	return slice
}
