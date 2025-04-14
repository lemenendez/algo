package ds

import "errors"

type SliceQueue[K comparable] struct {
	items []K
}

func (q *SliceQueue[K]) Enqueue(ele K) {
	// prepend
	// https://stackoverflow.com/questions/53737435/how-to-prepend-int-to-slice
	q.items = append([]K{ele}, q.items...)
}

func (q *SliceQueue[K]) Deque() (K, error) {
	if len(q.items) == 0 {
		return *(new(K)), errors.New("queue is empty")
	}
	index := len(q.items) - 1
	ele := q.items[index]
	q.items = q.items[:index]
	return ele, nil
}

func NewSliceQueue[K comparable]() *SliceQueue[K] {
	return &SliceQueue[K]{}
}
