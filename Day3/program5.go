package main

import "fmt"

func main() {

	var t float64 = 20.0
	fmt.Println("Tinggi tabung = 20")
	var r float64 = 4.0
	fmt.Println("Jari-jari tabung = 4")
	var pi float64 = 3.14
	fmt.Println("π = 3.14 (22/7) ")

	fmt.Println("Luas permukaan tabung yaitu =", (2*pi*r*t)+(2*pi*r*r))
}
