package main

import "fmt"

// arrInsertGeneric given two arrays and an index, merge the second one into the first one starting in the index
func arrInsertGeneric[k comparable](arr1 []k, arr2 []k, index int) []k {
	l1 := len(arr1)
	l2 := len(arr2)
	ret := make([]k, l2+l1)
	for i := 0; i < l1; i++ {
		if i < index {
			ret[i] = arr1[i]
			continue
		}
		ret[l2+i] = arr1[i]
	}

	for i := 0; i < l2; i++ {
		ret[index+i] = arr2[i]
	}

	return ret
}

// arrInsert given two arrays and an index, merge the second one with the first one starting in the given index
func arrInsert(arr1 []int, arr2 []int, index int) []int {
	l1 := len(arr1)
	l2 := len(arr2)
	ret := make([]int, l2+l1)
	for i := 0; i < l1; i++ {
		if i < index {
			ret[i] = arr1[i]
			continue
		}
		ret[l2+i] = arr1[i]
	}

	for i := 0; i < l2; i++ {
		ret[index+i] = arr2[i]
	}

	return ret
}

// insert from the book graphic go algorithms
func insert(arr []int, length int, tempArray []int, value int, insertIndex int) {
	for i := 0; i < length; i++ {
		if i < insertIndex {
			tempArray[i] = arr[i]
		} else {
			tempArray[i+1] = arr[i]
		}
	}
	tempArray[insertIndex] = value
}

// main func from the book graphic go algorithms
func originalMain() {
	var scores = []int{90, 70, 50, 80, 60, 85}
	var l = len(scores)
	var tempArray = make([]int, l+1)
	insert(scores, l, tempArray, 75, 2)
	scores = tempArray
	for i := 0; i < l+1; i++ {
		fmt.Printf("%d ", scores[i])
	}
}

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}

	var more = []int{666, 900}

	all := arrInsert(scores, more, 2)
	fmt.Println(all)

	allgen := arrInsertGeneric(scores, more, 2)
	fmt.Println(allgen)

	var phrase1 = []string{"The", "only", "thing", "is", "fear", "itself", "."}
	var phrase2 = []string{"we", "have", "to", "fear"}
	phrase := arrInsertGeneric(phrase1, phrase2, 3)
	fmt.Println(phrase)
}
