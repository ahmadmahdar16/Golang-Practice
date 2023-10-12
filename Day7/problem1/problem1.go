package main

import (
	"fmt"
)

func Compare(a, b string) string {
	sub := len(a)
	if len(b) < sub {
		sub = len(b)
	}
	for i := 0; i < sub; i++ {
		if a[i] != b[i] {
			return a[:i]
		}
	}
	return a[:sub]
}
func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     // AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KI", "KIJANG"))      // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
}
