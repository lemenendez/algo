package main

import (
	"fmt"
	"math"
	"sort"
)

/*
k closes points to origin
https://leetcode.com/problems/k-closest-points-to-origin/description/
*/

func kClosest(points [][]int, k int) [][]int {
	type Point struct {
		x, y int
		d    float64
	}
	Points := make([]Point, len(points))

	for i := 0; i < len(points); i++ {
		x, y := points[i][0], points[i][1]
		Points[i] = Point{x, y, math.Sqrt(float64(x*x) + float64(y*y))}
	}
	sort.Slice(Points, func(i, j int) bool {
		return Points[i].d < Points[j].d
	})
	var answer [][]int
	for i := 0; i < k; i++ {
		answer = append(answer, []int{Points[i].x, Points[i].y})
	}
	return answer
}

func main() {
	matrix := [][]int{{1, 3}, {-2, 2}}
	fmt.Println(kClosest(matrix, 1))
	fmt.Println(kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))

}
