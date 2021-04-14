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

//FibonnaciIterative finds the fibonnacci of a given index by iterating
func FibonnaciIterative(index int) int {
	if index < 2 {
		return index
	}
	num1 := 0
	num2 := 1
	var answer int
	for index > 1 {
		answer = num1 + num2
		num1, num2 = num2, answer
		index--
	}
	return answer
}

//FibonnaciRecursive finds the fibonnacci of a given index recursivly
func FibonnaciRecursive(index int) int {
	if index < 2 {
		return index
	}
	return FibonnaciRecursive(index-1) + FibonnaciRecursive(index-2)
}
