package algo

//FindFactorialRecursive finds the factorial of a given number recursivly
func FindFactorialRecursive(number int) int {
	if number == 1 {
		return 1
	}
	return number * FindFactorialRecursive(number-1)
}
