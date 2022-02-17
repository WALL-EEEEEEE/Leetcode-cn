package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (heap *IntHeap) Len() int {
	return len(*heap)
}
func (heap *IntHeap) Less(i int, j int) bool {
	return (*heap)[i] < (*heap)[j]
}

func (heap *IntHeap) Swap(i int, j int) {
	tmp := (*heap)[i]
	(*heap)[i] = (*heap)[j]
	(*heap)[j] = tmp
}

func (heap *IntHeap) Push(x interface{}) {
	*heap = append(*heap, x.(int))
}

func (heap *IntHeap) Pop() interface{} {
	v := (*heap)[len(*heap)-1]
	*heap = (*heap)[:len(*heap)-1]
	return v
}

type KthLargest struct {
	values IntHeap
	k      int
}

func (this *KthLargest) Add(val int) int {
	heap.Push(&this.values, val)
	if this.values.Len() > this.k {
		heap.Pop(&this.values)
	}
	return this.values[0]
}

func Constructor(k int, nums []int) KthLargest {
	kth_large := &KthLargest{
		k: k,
	}
	heap.Init(&kth_large.values)
	for i, v := range nums {
		heap.Push(&kth_large.values, v)
		if i > k-1 {
			heap.Pop(&kth_large.values)
		}
	}
	return *kth_large
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
func Kth_Largest(k int, nums []int, stream []int) {
	obj := Constructor(k, nums)
	for _, v := range stream {
		k := obj.Add(v)
		fmt.Println(k)
	}
}

func main() {
	test_cases := []interface{}{
		[]interface{}{
			[]interface{}{
				3,
				[]int{4, 5, 8, 2},
			},
			[]int{3, 5, 10, 9, 4},
		},
		[]interface{}{
			[]interface{}{
				1,
				[]int{},
			},
			[]int{-3, -2, -4, 0, 4},
		},
	}
	for _, test_case := range test_cases {
		k := test_case.([]interface{})[0].([]interface{})[0].(int)
		values := test_case.([]interface{})[0].([]interface{})[1].([]int)
		stream := test_case.([]interface{})[1].([]int)
		Kth_Largest(k, values, stream)
	}

}
