package piscine

func Any(f func(string) bool, a []string) bool {
	var b bool
	for i := 0; i < len(a); i++ {
		b = f(a[i])
		if f(a[i]) == true {
			return f(a[i])
		}
	}
	return b
}
