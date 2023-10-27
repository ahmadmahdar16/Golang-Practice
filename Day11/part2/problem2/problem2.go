package main

import "fmt"

func BottomUp(n int) int {
	if n <= 1 {
		return n
	}
	fib := make([]int, n+1)
	fib[1] = 1
	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib[n]
}

func main() {
	fmt.Println(BottomUp(0))  //0
	fmt.Println(BottomUp(1))  //1
	fmt.Println(BottomUp(2))  //1
	fmt.Println(BottomUp(3))  //2
	fmt.Println(BottomUp(5))  //5
	fmt.Println(BottomUp(6))  //8
	fmt.Println(BottomUp(7))  //13
	fmt.Println(BottomUp(9))  //34
	fmt.Println(BottomUp(10)) //55
}
