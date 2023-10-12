package main

import "fmt"

func ArrayUnique(arrayA, arrayB []int) []int {
	Unarr := make(map[int]bool)
	for _, unique := range arrayA {
		Unarr[unique] = true
	}

	for _, unique := range arrayB {
		if _, ok := Unarr[unique]; ok {
			delete(Unarr, unique)
		}
	}

	UnarrSlice := make([]int, 0, len(Unarr))
	for unique := range Unarr {
		UnarrSlice = append(UnarrSlice, unique)
	}
	return UnarrSlice
}

func main() {
	fmt.Println(ArrayUnique([]int{1, 2, 3, 4}, []int{1, 3, 5, 10, 16}))   //[2,4]
	fmt.Println(ArrayUnique([]int{10, 20, 30, 40}, []int{5, 10, 15, 59})) //[20, 30, 40]
	fmt.Println(ArrayUnique([]int{1, 3, 7}, []int{1, 3, 5}))              //[7]
	fmt.Println(ArrayUnique([]int{3, 8}, []int{2, 8}))                    //[3]
	fmt.Println(ArrayUnique([]int{1, 2, 3}, []int{3, 2, 1}))              //[]
}
