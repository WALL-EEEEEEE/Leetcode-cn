package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*



func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func partSort(arr []int, left, right int) int {
	k := rand.Intn(right - left)
	swap(arr, k, right)
	pivot := arr[right]
	i, j := left-1, left
	for j < right {
		if arr[j] < pivot {
			i++
			swap(arr, i, j)
		}
		j++
	}
	swap(arr, i+1, j)
	return i + 1
}

func findKthLargest(nums []int, k int) int {
	size := len(nums)
	var left, right, pivot int = 0, size - 1, size - 1
	k = size - k
	pivot = partSort(nums, left, right)
	for pivot <= right {
		if pivot > k {
			pivot = partSort(nums, left, pivot-1)
			continue
		} else if pivot < k {
			pivot = partSort(nums, pivot+1, right)
			continue
		}
		break
	}
	return nums[k]
}

*/

func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(a []int, l, r, index int) int {
	q := partition(a, l, r) //randomPartition(a, l, r)
	if q == index {
		return a[q]
	} else if q < index {
		return quickSelect(a, q+1, r, index)
	}
	return quickSelect(a, l, q-1, index)
}

func randomPartition(a []int, l, r int) int {
	var i int
	if r-l > 0 {
		i = rand.Intn(r-l) + l
	} else {
		i = l
	}
	a[i], a[r] = a[r], a[i]
	return partition(a, l, r)
}

func partition(a []int, l, r int) int {
	x := a[r]
	i := l - 1
	for j := l; j < r; j++ {
		if a[j] <= x {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
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
		[]interface{}{
			[]int{-1, 2, 0},
			3,
		},
		[]interface{}{
			[]int{3, 1, 2, 4},
			2,
		},
		[]interface{}{
			[]int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6},
			20,
		},
	}
	for _, test_case := range test_cases {

		nums := test_case.([]interface{})[0].([]int)
		k := test_case.([]interface{})[1].(int)
		fmt.Printf("数组：%v， ", nums)
		kth := findKthLargest(nums, k)
		fmt.Printf("第%v个最大值为：%v\n", k, kth)
	}

}
