package main

import (
	"container/heap"
	"fmt"
)

type IntMaxHeap []int

func (heap *IntMaxHeap) Len() int {
	return len(*heap)
}

func (heap *IntMaxHeap) Less(i int, j int) bool {
	return (*heap)[i] > (*heap)[j]
}

// Swap swaps the elements with indexes i and j.
func (heap *IntMaxHeap) Swap(i int, j int) {
	tmp := (*heap)[i]
	(*heap)[i] = (*heap)[j]
	(*heap)[j] = tmp
}

func (heap *IntMaxHeap) Push(x interface{}) {
	*heap = append(*heap, x.(int))
}

func (heap *IntMaxHeap) Pop() interface{} {
	v := (*heap)[len(*heap)-1]
	*heap = (*heap)[:len(*heap)-1]
	return v
}

func lastStoneWeight(stones []int) int {
	var maxheap IntMaxHeap = stones
	heap.Init(&maxheap)
	for maxheap.Len() > 1 {
		first := heap.Pop(&maxheap).(int)
		second := heap.Pop(&maxheap).(int)
		if first != second {
			heap.Push(&maxheap, first-second)
		}
	}
	if maxheap.Len() == 0 {
		return 0
	}
	return heap.Pop(&maxheap).(int)
}

func main() {
	test_cases := [][]int{
		[]int{2, 7, 4, 1, 8, 1},
	}
	for _, test_case := range test_cases {
		fmt.Printf("石头重量: %v, ", test_case)
		last_stone := lastStoneWeight(test_case)
		fmt.Printf("最后的石头重量: %v", last_stone)
	}
}
