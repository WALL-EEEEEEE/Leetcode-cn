package main

import "fmt"

func pardition(arr []int, left int, right int) int {
	pivot := arr[right]
	var i, j int
	for i, j = left-1, left; j < right; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}

func median_in_array(arr []int) int {
	var median_index int = len(arr) / 2
	left, right := 0, len(arr)-1
	p := pardition(arr, left, right)
	for p != median_index {
		if p < median_index {
			p = pardition(arr, left+1, right)
		} else if left > median_index {
			p = pardition(arr, 0, left-1)
		} else {
			break
		}
	}
	if len(arr)%2 == 0 {
		return (arr[p] + arr[p-1]) / 2
	}
	return arr[p]
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
