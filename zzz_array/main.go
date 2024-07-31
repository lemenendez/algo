// THIS IS INCOMPLETE FOR THE MOMENT
package main

import (
	"errors"
	"fmt"
	"math"
	"slices"
)

const (
	max_array_size = math.MaxInt
	notFound       = -1
)

var indexOutOfBoundError = errors.New("index out of bound")

type IntArray struct {
	data []int
	size int
}

func (arr *IntArray) rangeCheck(index int) error {
	if index < 0 || index > arr.size {
		return indexOutOfBoundError
	}
	return nil
}

func (arr *IntArray) Set(index int, value int) (int, error) {
	if err := arr.rangeCheck(index); err != nil {
		return notFound, err
	}
	old := arr.data[index]
	arr.data[index] = value
	return old, nil
}

func (arr *IntArray) Get(index int) (int, error) {
	if err := arr.rangeCheck(index); err != nil {
		return notFound, err
	}
	return arr.data[index], nil
}

func (arr *IntArray) Add(value int) error {
	arr.data = append(arr.data, value)
	arr.size++
	return nil
}

func (arr *IntArray) DeleteByIndex(index int) error {
	if err := arr.rangeCheck(index); err != nil {
		return err
	}
	slices.Delete(arr.data, index, index)
	return nil
}

func (arr *IntArray) Len() int { return arr.size }

func NewArray() *IntArray {
	return &IntArray{
		data: make([]int, 0),
	}
}

func main() {
	arr := NewArray()
	if err := arr.Add(3); err != nil {
		panic(err)
	}
	for i := 0; i < arr.Len(); i++ {
		val, err := arr.Get(i)
		if err != nil {
			panic(err)
		}
		fmt.Print(val)
	}
}
