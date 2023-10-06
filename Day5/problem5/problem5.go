package main

import (
	"fmt"
)

func pangkat(base, exponent int) int {
	var isNegative bool
	var output int
	output = 1

	if exponent < 0 {
		isNegative = true
		exponent = exponent * -1
	}

	for exponent != 0 {
		output *= base
		exponent -= 1
	}

	if isNegative == true {
		output = 1 / output
	}

	return output
}

func main() {

	fmt.Println(pangkat(2, 3))
	fmt.Println(pangkat(5, 3))
	fmt.Println(pangkat(10, 2))
	fmt.Println(pangkat(2, 5))
	fmt.Println(pangkat(7, 3))
}
