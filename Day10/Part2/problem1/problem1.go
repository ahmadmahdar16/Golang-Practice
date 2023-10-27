package main

import (
	"fmt"
)

func FindMinAndMax(arr []int) string {
	if len(arr) == 0 {
		return ""
	}

	min, max := arr[0], arr[0]
	minIndex, maxIndex := 0, 0

	for i, value := range arr {
		if value < min {
			min = value
			minIndex = i
		} else if value > max {
			max = value
			maxIndex = i
		}
	}

	return fmt.Sprintf("min: %d index: %d max: %d index: %d", min, minIndex, max, maxIndex)
}

func main() {
	fmt.Println(FindMinAndMax([]int{5, 7, 4, -2, -1, 8}))
	fmt.Println(FindMinAndMax([]int{2, -5, -4, 22, 7, 7}))
	fmt.Println(FindMinAndMax([]int{4, 3, 9, 4, -21, 7}))
	fmt.Println(FindMinAndMax([]int{-1, 5, 6, 4, 2, 18}))
	fmt.Println(FindMinAndMax([]int{-2, 5, -7, 4, 7, -20}))
}
