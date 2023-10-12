package main

import "fmt"

func primeNumber(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("11 adalah", primeNumber(11))
	fmt.Println("13 adalah", primeNumber(13))
	fmt.Println("17 adalah", primeNumber(17))
	fmt.Println("20 adalah", primeNumber(20))
	fmt.Println("35 adalah", primeNumber(35))
}
