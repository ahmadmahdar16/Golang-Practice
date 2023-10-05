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

	var kata string
	fmt.Scanln(&kata)

	for i := 0; i < len(kata)/2; i++ {
		if string(kata[i]) != string(kata[len(kata)-1-i]) {
			fmt.Println("False")
			return
		}
	}
	fmt.Println("True")

}
