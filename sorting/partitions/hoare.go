package main

import "fmt"

/*

Partitioning is a technique used in sorting algorithms like quicksort to divide an array into two parts based on a pivot element.
The goal of partitioning is to rearrange the elements in such a way that all elements less than the pivot come before it,
and all elements greater than the pivot come after it.
The partitioning process is crucial for the efficiency of quicksort, as it allows the algorithm to work on smaller subarrays recursively.
The partitioning process is typically done in place, meaning that it doesn't require additional memory allocation for a new array.
Instead, it rearranges the elements within the original array.

We have to reorder elements of the array such that in the result array:

- Pivot is the correct index as per sorting order
- All elements before pivot should be less thant pivot but not necessarily in order
- All elements after pivot should be larger thant pivot but not necessarily in sorted order.

https://www.geeksforgeeks.org/hoare-s-partition-algorithm/

*/

func hoare(arr []int, l, h int) int {
	if arr == nil {
		return -1
	}
	//l := 0
	//h := len(arr) - 1
	pivot := arr[l]
	for {
		for arr[l] < pivot {
			l++
		}
		for arr[h] > pivot {
			h--
		}
		if l >= h {
			return h
		}
		arr[l], arr[h] = arr[h], arr[l]
	}
}

func main() {
	cases := []struct {
		arr []int
		l   int
		h   int
		k   int
	}{
		{arr: []int{1, 2}, l: 0, h: 1, k: 0},
		{arr: []int{2, 1}, l: 0, h: 1, k: 1},
		{arr: []int{5, 3, 8, 4, 2, 7, 1, 10}, l: 0, h: 7, k: 4},
		{arr: []int{12, 10, 9, 16, 19, 9}, l: 0, h: 5, k: 3},
		{arr: []int{4, 10, 9, 8, 16, 19, 9}, l: 0, h: 6, k: 0},
	}
	for i, c := range cases {
		k := hoare(c.arr, c.l, c.h)
		if k != c.k {
			fmt.Printf(" id %v arr %v unexpected result %v, expected %v\n", i, c.arr, k, c.k)
		}
	}
}
