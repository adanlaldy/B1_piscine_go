package piscine

func UltimateDivMod(a *int, b *int) {
	ab := *a / *b
	*b = *a % *b
	*a = ab
}
