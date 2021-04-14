package main

import (
	"fmt"

	"github.com/louissaadgo/ds-algo/algo"
)

func main() {
	answer := algo.FindFactorialRecursive(10)
	fmt.Println(answer)
	answer2 := algo.FindFactorialIterative(10)
	fmt.Println(answer2)
	answer3 := algo.FibonnaciIterative(10)
	fmt.Println(answer3)
	answer4 := algo.FibonnaciRecursive(10)
	fmt.Println(answer4)
	answer5 := algo.ReverseStringIterative("Hi Louis")
	fmt.Println(answer5)
}
