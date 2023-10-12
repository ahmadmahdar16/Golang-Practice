package main

import "fmt"

func palindrome(input string) bool {
	for i := 1; i < len(input)/2; i++ {
		if string(input[i]) != string(input[len(input)-1-i]) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(palindrome("civic"))
	fmt.Println(palindrome("katak"))
	fmt.Println(palindrome("kasur rusak"))
	fmt.Println(palindrome("kupu-kupu"))
	fmt.Println(palindrome("lion"))
}
