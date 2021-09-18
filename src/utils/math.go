package math

// This shows the basic synatc for a functoin
func Factorial(num int) (val int) {
	if num <= 1 {
		return 1
	}
	return num * Factorial(num-1)
}
