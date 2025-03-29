package main

import (
	"cmp"
	"fmt"
)

func bubbleSortGeneric[k cmp.Ordered](arr []k) []k {
	for i := 0; i < len(arr)-1; i++ {
		swapped := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if swapped == false {
			break
		}
	}
	return arr
}

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		swapped := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if swapped == false {
			break
		}
	}
	return arr
}

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}
	scores = bubbleSort(scores)
	fmt.Print(scores)

	names := bubbleSortGeneric([]string{"adios", "mundo", "cruel"})
	fmt.Println(names)
}
