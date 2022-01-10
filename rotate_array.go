package main

import "fmt"

func rotate(nums []int, k int) {
	arr_len := len(nums)
	k = k % arr_len
	var tmp []int = make([]int, arr_len-k)
	copy(tmp[:], nums[:arr_len-k])
	copy(nums[:k], nums[arr_len-k:])
	copy(nums[k:], tmp)
}

type Pair struct {
	array       []int
	rotate_step int
}

func main() {
	tests := []Pair{
		Pair{
			array:       []int{1, 2, 3, 4, 5, 6, 7},
			rotate_step: 4,
		},
		Pair{
			array:       []int{-1, -2},
			rotate_step: 3,
		},
	}
	for _, test := range tests {

		fmt.Printf("数组：%v ,", test.array)
		rotate(test.array, test.rotate_step)
		fmt.Printf("轮转%d次后：%v \n", test.rotate_step, test.array)
	}
}
