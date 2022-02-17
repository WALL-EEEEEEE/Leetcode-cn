package main

import (
	"container/heap"
	"fmt"
)

type minHeap []int

func (h *minHeap) Len() int {
	return len(*h)
}

func (h *minHeap) Less(i int, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *minHeap) Swap(i int, j int) {
	tmp := (*h)[i]
	(*h)[i] = (*h)[j]
	(*h)[j] = tmp
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	size := h.Len()
	v := (*h)[size-1]
	*h = (*h)[:size-1]
	return v
}

func findKthLargest(nums []int, k int) int {
	var h minHeap = nums
	heap.Init(&h)
	for h.Len() > k {
		heap.Pop(&h)
	}
	return h[0]
}

func main() {
	test_cases := []interface{}{
		[]interface{}{
			[]int{3, 2, 1, 5, 6, 4},
			2,
		},
		[]interface{}{
			[]int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			4,
		},
	}
	for _, test_case := range test_cases {
		nums := test_case.([]interface{})[0].([]int)
		k := test_case.([]interface{})[1].(int)
		kth := findKthLargest(nums, k)
		fmt.Printf("数组：%v， 第%v个最大值为：%v\n", nums, k, kth)
	}

}
