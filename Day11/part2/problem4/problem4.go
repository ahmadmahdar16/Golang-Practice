package main

import "fmt"

func RomanNumerals(value int) string {
	//your code here
	if value <= 0 || value > 3999 {
		return ""
	}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman := ""
	for i := 0; i < len(values); i++ {
		for value >= values[i] {
			roman += symbols[i]
			value -= values[i]
		}
	}
	return roman
}

func main() {
	fmt.Println(RomanNumerals(6))    //VI
	fmt.Println(RomanNumerals(9))    //IX
	fmt.Println(RomanNumerals(23))   //XXIII
	fmt.Println(RomanNumerals(2021)) //MMXXI
	fmt.Println(RomanNumerals(1646)) //MDCXLVI
}
