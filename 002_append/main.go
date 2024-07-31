package main

import "fmt"

// customAppendGeneric given an array of ints and an element, it appends and element at the end, then returns the array
func customAppendGeneric[k comparable](arr []k, elem k) []k {
	var length = len(arr)
	var tmp = make([]k, length+1)
	for i := 0; i < length; i++ {
		tmp[i] = arr[i]
	}
	tmp[length] = elem
	return tmp
}

// customAppend given an array of ints and an element, it appends and element at the end, then returns the array
func customAppend(arr []int, elem int) []int {
	var length = len(arr)
	var tmp = make([]int, length+1)
	for i := 0; i < length; i++ {
		tmp[i] = arr[i]
	}
	tmp[length] = elem
	return tmp
}

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}
	scores = customAppend(scores, 75)
	fmt.Println(scores)

	scores = customAppendGeneric(scores, 110)
	fmt.Println(scores)

}
