package main

import "fmt"

func FindaMaxSumSubArray(k int, arr []int) int {
	maxSum := 0
	for i := 0; i < len(arr)-k+1; i++ {
		sum := 0
		for j := 0; j < k; j++ {
			sum += arr[i+j]
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}

func main() {
	fmt.Println(FindaMaxSumSubArray(3, []int{2, 1, 5, 1, 3, 2})) //9
	fmt.Println(FindaMaxSumSubArray(2, []int{2, 3, 4, 1, 5}))    //7
	fmt.Println(FindaMaxSumSubArray(2, []int{2, 1, 4, 1, 1}))    //5
	fmt.Println(FindaMaxSumSubArray(3, []int{2, 1, 4, 1, 1}))    //7
	fmt.Println(FindaMaxSumSubArray(4, []int{2, 1, 4, 1, 1}))    //8
}
