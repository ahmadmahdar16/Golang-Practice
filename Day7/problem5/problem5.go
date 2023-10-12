package main

import "fmt"

func RemoveDuplicates(array []int) int {
	duplicate := make(map[int]bool)

	for _, dupe := range array {
		duplicate[dupe] = true
	}

	duplicateSlice := make([]int, 0, len(duplicate))
	for dupe := range duplicate {
		duplicateSlice = append(duplicateSlice, dupe)
	}
	return len(duplicateSlice)
}

func main() {
	fmt.Println(RemoveDuplicates([]int{2, 3, 3, 3, 6, 9, 9})) //4
	fmt.Println(RemoveDuplicates([]int{2, 3, 4, 5, 6, 9, 9})) //6
	fmt.Println(RemoveDuplicates([]int{2, 2, 2, 11}))         //2
	fmt.Println(RemoveDuplicates([]int{2, 2, 2, 11}))         //2
	fmt.Println(RemoveDuplicates([]int{1, 2, 3, 11, 11}))     //4
}
