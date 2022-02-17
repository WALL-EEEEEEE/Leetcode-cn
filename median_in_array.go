package main

import "fmt"

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func part_sort(arr []int, left int, right int) int {
	key := arr[right]
	i, j := left, left
	for j < right {
		if arr[j] < key {
			i++
			swap(arr, i, j)
		}
		j++
	}
	swap(arr, i+1, j)
	return i + 1
}

func median_in_array(arr []int) int {
	var median_index int
	median_index = len(arr) / 2
	q, left, right := 0, 0, len(arr)-1
	q = part_sort(arr, left, right)
	for q <= right {
		if q < median_index {
			q = part_sort(arr, q+1, right)
		} else if q > median_index {
			q = part_sort(arr, left, q-1)
		} else {
			break
		}
	}
	if len(arr)%2 == 0 {
		return (arr[q] + arr[q-1]) / 2
	}
	return arr[q]
}

func main() {
	test_cases := [][]int{
		[]int{
			1, 3, 5, 20, 0,
		},
		[]int{
			7, 9, 5, 20, 0, 6,
		},
		[]int{
			2, 1, 0, 4, 3,
		},
	}
	for _, test_case := range test_cases {
		fmt.Printf("数组: %v, ", test_case)
		median := median_in_array(test_case)
		fmt.Printf("中位数为：%v\n", median)
	}
}
