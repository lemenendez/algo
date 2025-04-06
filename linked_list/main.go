package main

import (
	"fmt"
	"github.com/lemenendez/algo/ds"
)

func main() {

	ls2 := ds.NewLinkedList(1)
	ls2.Append(3).
		Append(4).
		Append(2)
	ls2.Walker(func(k int) {
		fmt.Println(k)
	})

	mlst := ds.NewLinkedList(1).
		Append(4).
		Append(6).
		Merge(ds.NewLinkedList(1).Append(3).Append(5))
	mlst.Walker(func(k int) {
		fmt.Println(k)
	})

}
