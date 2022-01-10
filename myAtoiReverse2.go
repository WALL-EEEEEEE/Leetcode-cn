package main

import (
	"fmt"
)

func myAtoiReverse(s string) string {
	target := ""
	for i := len(s) - 1; i > 0; i-- {
		if (s[i]-'0' < 0 || s[i]-'0' > 9) && s[i] != '.' && s[i] != '-' {
			continue
		}
		target += string(s[i])
	}
	return target
}

func main() {
	num_strs := []string{
		"abc12.3",
		"abc-12.3",
		"abc12.3-",
		"abc12.-3",
	}
	for _, num_str := range num_strs {
		num := myAtoiReverse(num_str)
		fmt.Printf("%s 转换: %s\n", num_str, num)
	}
}
