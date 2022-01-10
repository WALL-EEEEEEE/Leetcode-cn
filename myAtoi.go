package main

import (
	"fmt"
	"math"
)

func myAtoi(str string) int {
	var result float64
	var flag int = 1
	var numbered bool = false
	var signed bool = false
	for _, v := range str {
		ivalue := int64(v - '0')
		if (ivalue < 0 || ivalue > 10) && v != '-' && v != ' ' && v != '+' {
			break
		} else if v == ' ' {
			if signed || numbered {
				break
			}
			continue
		} else if v == '-' {
			if signed || numbered {
				break
			}
			flag = -1
			signed = true
			continue
		} else if v == '+' {
			if signed || numbered {
				break
			}
			signed = true
			flag = 1
			continue
		}
		result = result*10 + float64(ivalue)
		numbered = true
	}
	result = result * float64(flag)
	if flag > 0 {
		result = float64(math.Min(math.MaxInt32, float64(result)))
	} else {
		result = float64(math.Max(math.MinInt32, float64(result)))
	}
	return int(result)
}

func main() {
	nums_str := []string{
		"110",
		"120",
		"-220",
		"-110",
		"-42",
		" -42",
		"4193 with words",
		"words and 987",
		"-91283472332",
		"+1",
		"-+1",
		"+-1",
		"+-12",
		"42",
		"00000-42a1234",
		"   +0 123",
		"9223372036854775808",
	}
	for _, num_str := range nums_str {
		num := myAtoi(num_str)
		fmt.Printf("%s 转成: %d\n", num_str, num)
	}
}
