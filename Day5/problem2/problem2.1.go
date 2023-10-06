package main

import "fmt"

func cariFaktor(bilangan int) []int {
	faktor := []int{}
	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			faktor = append(faktor, i)
		}
	}
	return faktor
}

func main() {
	var bilangan int
	fmt.Println("Masukkan bilangan")
	fmt.Scanf("%d", &bilangan)
	faktor := cariFaktor(bilangan)
	fmt.Println("faktor dari", bilangan, "adalah", faktor)

}
