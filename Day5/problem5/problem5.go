package main

import "fmt"

func pangkat(base, pangkat int) int {
	for i := 0; i < base; i++ {
		if base%pangkat == 0 {
		}
	}
	return 0
}

func main() {
	var x int
	var j int
	var res = 1
	fmt.Scanln(&x)
	fmt.Scanln(&j)

	for i := 0; i < x; i++ {
		res *= 2
	}
	fmt.Println(res)
}
