package main

import "fmt"

func Swap(a, b *int) (int, int) {
	//your code here
	*a, *b = *b, *a
	return *a, *b
}

func main() {
	a := 10
	b := 20

	fmt.Println("Input :", a, b)
	Swap(&a, &b)
	fmt.Println("Output:", a, b)
}
