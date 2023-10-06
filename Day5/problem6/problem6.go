package main

import "fmt"

func isPrima(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func FullPrima(n int) bool {
	for n > 0 {
		var single = n

		if n/10 > 0 {
			single = n % 10
		}
		if !isPrima(single) {
			return false
		}

		n /= 10
	}
	return true
}

func main() {
	fmt.Println(FullPrima(2))
	fmt.Println(FullPrima(3))
	fmt.Println(FullPrima(11))
	fmt.Println(FullPrima(13))
	fmt.Println(FullPrima(23))
	fmt.Println(FullPrima(29))
	fmt.Println(FullPrima(37))
	fmt.Println(FullPrima(41))
	fmt.Println(FullPrima(43))
	fmt.Println(FullPrima(53))
}
