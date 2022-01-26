package main

import "fmt"

func part_sort(arr []int, left int, right int) int {
	key := arr[right]
	for left < right {
		for left < right && key > arr[left] {
			left++
		}
		for left < right && key < arr[right] {
			right--
		}
		if left < right {
			tmp := arr[left]
			arr[left] = arr[right]
			arr[right] = tmp
		}
	}
	return left
}

func median_in_array(arr []int) int {
	var median_index int
	if len(arr)%2 == 0 {
		median_index = len(arr)/2 - 1
	} else {
		median_index = len(arr) / 2
	}
	left, right := 0, len(arr)-1
	for left < right {
		if left < median_index {
			left = part_sort(arr, left+1, right)
		} else if left > median_index {
			left = part_sort(arr, 0, left-1)
		} else {
			break
		}
	}
	if len(arr)%2 == 0 {
		return (arr[left] + arr[left+1]) / 2
	}
	return arr[left]
}

func main() {
	test_cases := [][]int{
		[]int{
			1, 3, 5, 20, 0,
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
