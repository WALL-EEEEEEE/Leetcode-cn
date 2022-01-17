package main

import "fmt"

func lexicalOrder(n int) []int {
	var ans []int = []int{}
	num := 1
	for i := 0; i < n; i = len(ans) {
		for num <= n {
			ans = append(ans, num)
			num *= 10
		}
		for num%10 == 9 || num > n {
			num /= 10
		}
		num += 1
	}
	return ans
}

func main() {
	var nums []int = []int{13, 10, 23}
	for _, num := range nums {
		lexical_nums := lexicalOrder(num)
		fmt.Printf("字典序遍历 %d: %+v\n", num, lexical_nums)
	}
}
