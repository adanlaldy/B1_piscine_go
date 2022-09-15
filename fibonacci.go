package piscine

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	} else if index > 0 {
		test := -Fibonacci(index-1) + Fibonacci(index-2)
		return test
	}
	return index
}
