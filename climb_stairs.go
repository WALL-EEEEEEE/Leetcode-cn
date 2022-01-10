package main

import "fmt"

/* s[n]为第n阶台阶的所爬的台阶数
 * 可以根据s[n]写出对应的状态方程：
 * 1. s[n] 等于 2, f(n) = f(n-2)
 * 2. s[n] 等于1， f(n) = f(n-1)
 * 那么状态方程为：f(n) = f(n-2) + f(n-1)
 */

func climbStairs(n int) int {
	f_n, f_n1, f_n2 := 0, 1, 0
	if n < 1 {
		return 1
	}
	for i := 0; i < n; i++ {
		f_n = f_n1 + f_n2
		f_n1, f_n2 = f_n, f_n1
	}
	return f_n
}

func main() {
	stairs := []int{
		0,
		1,
		2,
		3,
	}
	for _, stair := range stairs {
		solution := climbStairs(stair)
		fmt.Printf("台阶:%d阶，方式：%d\n", stair, solution)
	}
}
