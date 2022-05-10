package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func pardition(nums []string, left int, right int) int {
	size := right - left + 1
	pivot := rand.Intn(size) + left
	nums[pivot], nums[right] = nums[right], nums[pivot]
	j := left
	for i := left; i < right; i++ {
		n_i, _ := strconv.ParseInt(nums[i], 10, 64)
		n_r, _ := strconv.ParseInt(nums[right], 10, 64)
		if n_i-n_r <= 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	nums[j], nums[right] = nums[right], nums[j]
	return j
}

func quickSelect(nums []string, left int, right int, k int) string {
	pivot := pardition(nums, left, right)
	if pivot == k {
		return nums[pivot]
	}
	if pivot < k {
		return quickSelect(nums, pivot+1, right, k)
	}
	return quickSelect(nums, left, pivot-1, k)
}

func kthLargestNumber(nums []string, k int) string {
	k = len(nums) - k
	return quickSelect(nums, 0, len(nums)-1, k)

}

func main() {
	test_cases := []interface{}{
		/*
			[]interface{}{
				[]string{"3", "6", "7", "10"},
				4,
			},
			[]interface{}{
				[]string{"2", "21", "12", "1"},
				3,
			},
		*/
		[]interface{}{
			[]string{"683339452288515879", "7846081062003424420", "4805719838", "4840666580043", "83598933472122816064", "522940572025909479", "615832818268861533", "65439878015", "499305616484085", "97704358112880133", "23861207501102", "919346676", "60618091901581", "5914766072", "426842450882100996", "914353682223943129", "97", "241413975523149135", "8594929955620533", "55257775478129", "528", "5110809", "7930848872563942788", "758", "4", "38272299275037314530", "9567700", "28449892665", "2846386557790827231", "53222591365177739", "703029", "3280920242869904137", "87236929298425799136", "3103886291279"},
			3,
		},
	}
	for _, test_case := range test_cases {
		nums := test_case.([]interface{})[0].([]string)
		k := test_case.([]interface{})[1].(int)
		kth := kthLargestNumber(nums, k)
		fmt.Printf("数组：%v， 第%v个整数为：%v\n", nums, k, kth)
	}

}
