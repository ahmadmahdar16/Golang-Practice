package main

import "fmt"

func playWithAsterik(n int) {
	k := 2*n - 2
	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			fmt.Print(" ")
		}
		k = k - 1
		for j := 0; j < i+1; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func main() {
	playWithAsterik(5)
}
