package main

import "fmt"

/*
Partition is the problem where we are given an element (
PIVOT) and we have to reorder elements of the array such that in the result array:

- Pivot is the correct index as per sorting order
- All elements before pivot should be less thant pivot but not necessarily in order
- All elements after pivot should be larger thant pivot but not necessarily in sorted order.

*/

// hoare, pivot element is always the first element
// 1 Initialize two pointers L and H, L is the smallest index, H is the highest index
// 2 The pivot element at the smallest index
// 3.1 Keep increasing L while array[i] < pivot
// 3.2 Keep decreasing H while arr[j] > pivot
// 3.3 if i>=j, return j
// 3.4 if i<j, swap arr[i] and arr[j]
// [5,3,8,4,2,7,1,10]
func hoare(arr []int, l, h int) int {
	// initially l = 0, h = size -1
	pivot := arr[l]
	i := l
	j := h - 1
	for {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i >= j {
			return j
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	return -1
}

func main() {
	fmt.Println(hoare([]int{5, 3, 8, 4, 2, 7, 1, 10}, 0, 8))
}
