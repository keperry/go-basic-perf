package perf

func Factorial(n int) int {

	var result int
	if n == 0 || n == 1 {
		return 1
	}
	result = Factorial(n-1) * n
	return result
}
