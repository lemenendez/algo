package main

import "fmt"

// linearSearchGeneric given an array (haystack) of 'comparable' type and a needle of the same type
// returns the index of value present in the array or -1 if not found
func linearSearchGeneric[K comparable](arr []K, item K) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == item {
			return i
		}
	}
	return -1
}

// linearSearch2 given an array (haystack) and an integer (needle),
// returns all indexes of the needle present in the array
func linearSearch2(arr []int, target int) []int {
	result := []int{}
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			result = append(result, i)
		}
	}
	return result
}

// linearSearch given an array(haystack) and an integer (needle), returns the index of the needle present in the array
// returns -1 if not found
// complexity O(N)
func linearSearch(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(linearSearch([]int{1, 3, 5, 6}, 2))
	fmt.Println(linearSearch([]int{1, 3, 5, 6}, 5))

	fmt.Println(linearSearch2([]int{1, 3, 5, 6, 5}, 5))

	fmt.Println(linearSearchGeneric([]int{1, 3, 5, 6}, 2))
	fmt.Println(linearSearchGeneric([]int{1, 3, 5, 6}, 5))

}
