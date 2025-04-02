package main

import "fmt"

/*
Given an array of integers, return indices of the two numbers such that they add up to specific target.
You may assume that each input would have exactly one solution, and you many not use the same element twice.

Input: [2, 7, 11,15]
target = 9

Output: [0,1] because arr[0] + arr[1] = 2 + 7 = 9

https://leetcode.com/problems/two-sum/description/

*/

func find(arr []int, target int) (a, b int) {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		value := arr[i]
		rest := target - value
		if _, ok := m[rest]; ok {
			return i, m[rest]
		} else {
			m[value] = i
		}
	}
	return -1, -1
}

func main() {
	fmt.Println(find([]int{2, 7, 11, 15}, 9))     // 1,0
	fmt.Println(find([]int{2, 7, 11, 15}, 16000)) // -1,-1
	fmt.Println(find([]int{3, 2, 4}, 6))          // [1,2]
}
