package main

import (
	"fmt"
)

func numDecodings(s string) int {
	size := len(s)
	if size <= 0 {
		return 1
	}
	s_i1, s_i, s_i2 := 1, 0, 1
	for i := 1; i <= size; i++ {
		s_i = 0
		if s[i-1] != '0' {
			s_i += s_i1
		}
		if i > 1 && s[i-2] != '0' && (s[i-2]-'0')*10+s[i-1]-'0' <= 26 {
			s_i += s_i2
		}
		s_i1, s_i2 = s_i, s_i1
	}
	return s_i
}

func main() {
	encodes := []string{
		"12",
		"226",
		"06",
		"10",
		"27",
		"26",
		"2101",
	}
	for _, encode_str := range encodes {
		num := numDecodings(encode_str)
		fmt.Println("编码字符： ", encode_str, ", 解码：", num)
	}
}
