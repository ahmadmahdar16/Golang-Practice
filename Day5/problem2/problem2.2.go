package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scanf("%d", &bilangan)

	for i := bilangan; i >= 1; i-- {
		if bilangan%i == 0 {
			fmt.Println(i)
		}
	}
}
