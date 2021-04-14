package algo

//FindFactorialRecursive finds the factorial of a given number recursivly
func FindFactorialRecursive(number int) int {
	if number == 1 {
		return 1
	}
	return number * FindFactorialRecursive(number-1)
}

//FindFactorialIterative finds the factorial of a given number by iterating
func FindFactorialIterative(number int) int {
	answer := number
	for number > 1 {
		number--
		answer = answer * number
	}
	return answer
}
