package math

func Factorial(num int) (val int) {
	if num <= 1 {
		return 1
	}
	return num * Factorial(num-1)
}
