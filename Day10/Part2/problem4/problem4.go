package main

import "fmt"

func MostAppearItem(items []string) string {
	counts := make(map[string]int)

	for _, item := range items {
		counts[item]++
	}

	var result string
	for item, count := range counts {
		result += fmt.Sprintf("%s -> %d\n", item, count)
	}

	return result
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
}
