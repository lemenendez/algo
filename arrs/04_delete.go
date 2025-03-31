package main

import "fmt"

func removeGeneric[k comparable](arr []k, index int) []k {
	l := len(arr)
	res := make([]k, l-1)
	for i := 0; i < l; i++ {
		if i < index {
			res[i] = arr[i]
			continue
		}
		if i > index {
			res[i-1] = arr[i]
			continue
		}
	}
	return res
}

func remove(arr []int, index int) []int {
	l := len(arr)
	res := make([]int, l-1)
	for i := 0; i < l; i++ {
		if i < index {
			res[i] = arr[i]
			continue
		}
		if i > index {
			res[i-1] = arr[i]
			continue
		}
	}
	return res
}

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}
	scores = removeGeneric(scores, 2)
	fmt.Println(scores)
}
