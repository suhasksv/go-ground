package factorial

func Factorial(num int) int {
	if num == 1 {
		return num
	} else {
		return num * Factorial(num-1)
	}
}
