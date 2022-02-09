package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (heap IntHeap) Len() int {
	return len(heap)
}

func (heap IntHeap) Less(i, j int) bool {
	return heap[i] < heap[j]
}

func (heap *IntHeap) Pop() interface{} {
	old := *heap
	n := len(old)
	if n < 1 {
		return nil
	}
	v := old[n-1]
	*heap = old[0 : n-1]
	return v

}

func (heap *IntHeap) Push(i interface{}) {
	*heap = append(*heap, i.(int))
}

func (heap IntHeap) Swap(i, j int) {
	tmp := heap[i]
	heap[i] = heap[j]
	heap[j] = tmp
}

func median_in_array(arr []int) int {
	var len int = len(arr)
	var heap_len int
	if len%2 != 0 {
		heap_len = (len + 1) / 2
	} else {
		heap_len = len/2 + 1
	}
	var int_heap IntHeap = make([]int, 0, heap_len)
	heap.Init(&int_heap)
	for i, v := range arr {
		if i < heap_len {
			heap.Push(&int_heap, v)
			continue
		}
		if int_heap[0] < v {
			heap.Pop(&int_heap)
			heap.Push(&int_heap, v)
		}
	}
	if len%2 != 0 {
		return heap.Pop(&int_heap).(int)
	}
	return (heap.Pop(&int_heap).(int) + heap.Pop(&int_heap).(int)) / 2
}

func main() {
	test_cases := [][]int{
		[]int{
			5, 3, 1, 20, 0,
		},
		[]int{
			7, 9, 5, 20, 0, 6,
		},
	}
	for _, test_case := range test_cases {
		median := median_in_array(test_case)
		fmt.Printf("数组: %v, 中位数为：%v\n", test_case, median)
	}
}
