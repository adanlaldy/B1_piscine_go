package piscine

func Any(f func(string) bool, a []string) bool {
	var b bool
	for i := 0; i < len(a); i++ {
		b = IsNumeric(a[i])
		if IsNumeric(a[i]) == true {
			return IsNumeric(a[i])
		}
	}
	return b
}
