package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int //定义一个类型

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func main() {
	h := &IntHeap{
		2, 1, 5, 6, 4, 3, 7, 9, 8, 0,
	}
	heap.Init(h)
	fmt.Println(heap.Pop(h))
	heap.Push(h, 6)
	fmt.Println(*h)
	for len(*h) > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
