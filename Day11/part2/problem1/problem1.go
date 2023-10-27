package main

import "fmt"

func TopDown(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}
	if val, ok := memo[n]; ok {
		return val
	}
	memo[n] = TopDown(n-1, memo) + TopDown(n-2, memo)
	return memo[n]
}

func fibo(n int) int {
	memo := make(map[int]int)
	return TopDown(n, memo)
}

func main() {
	fmt.Println(fibo(0))  //0
	fmt.Println(fibo(1))  //1
	fmt.Println(fibo(2))  //1
	fmt.Println(fibo(3))  //2
	fmt.Println(fibo(5))  //5
	fmt.Println(fibo(6))  //8
	fmt.Println(fibo(7))  //13
	fmt.Println(fibo(9))  //34
	fmt.Println(fibo(10)) //55
}
