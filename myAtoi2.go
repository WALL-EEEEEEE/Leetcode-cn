package main

import (
	. "algorithm/structs"
	"fmt"
)

/*
	if this.getState() == "start" && (value == '-' || value == '+' || value == ' ') {
		if value == ' ' {
			return
		}
		if value == '-' {
			this.sign = -1
		}
		this.setState("signed")
	} else if (this.getState() == "numbered" || this.getState() == "signed" || this.getState() == "start") && unicode.IsNumber(value) {
		if this.result.(int) > math.MaxInt32/10 || (this.result.(int) == math.MaxInt32/10 && math.MaxInt32%10 < int(value-'0')) {
			if this.sign < 0 {
				this.result = math.MaxInt32 + 1
			} else {
				this.result = math.MaxInt32
			}
		} else {
			this.result = this.result.(int)*10 + int(value-'0')
		}
		this.setState("numbered")
	} else {
		this.setState("end")
	}

*/
func myAtoi(str string) int {

	dfm := NewDFM("start", map[string]string{"start": "", "signed": "", "numbered": "", "end": ""}, 0, 1)
	for _, v := range str {
		if dfm.End() {
			break
		}
		dfm.Step(v)
	}
	return dfm.Result().(int) * dfm.Sign()
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
		"21474836460",
	}
	for _, num_str := range nums_str {
		num := myAtoi(num_str)
		fmt.Printf("%s 转成: %d\n", num_str, num)
	}
}
