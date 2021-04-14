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
}
