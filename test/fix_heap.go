package main

import (
	"container/heap"
	"fmt"
	"math/rand"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

const (
	n = 10
	m = 3
)

func main() {
	h := &IntHeap{}
	heap.Init(h)

	for i := 0; i < n; i++ {
		x := rand.Intn(n) + 1
		fmt.Printf("%v, ", x)
		heap.Push(h, x)
		if h.Len() > m {
			heap.Pop(h)
		}
	}
	fmt.Println()

	r := make([]int, h.Len())
	for i := len(r) - 1; i >= 0; i-- {
		r[i] = heap.Pop(h).(int)
	}
	fmt.Printf("%v\n", r)
}
