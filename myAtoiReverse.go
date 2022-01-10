package main

import (
	. "algorithm/structs"
	"fmt"
	"unicode"
)

const (
	start State = iota
	sign
	number
	float
	end
)

func valid_char(value rune) bool {
	if value != '+' && value != '-' && value != ' ' && value != '.' && !unicode.IsNumber(value) {
		return false
	}
	return true
}

func from_start(this *DFM, value interface{}) State {
	if !valid_char(value.(rune)) {
		return end
	}
	if value.(rune) == ' ' {
		return start
	}
	if unicode.IsNumber(value.(rune)) {
		return from_number(this, value)
	}
	return end
}

func from_float(this *DFM, value interface{}) State {
	if !unicode.IsNumber(value.(rune)) && value.(rune) != '-' && value.(rune) != '+' {
		return end
	}
	if value.(rune) == '-' {
		this.SetData("sign", float32(-1))
		return sign
	} else if value.(rune) == '+' {
		return sign
	}
	result := this.GetData("result").(float32)
	float_base := this.GetData("float_base").(float32) * 0.1
	result = result + float32(value.(rune)-'0')*float_base
	this.SetData("float_base", float_base)
	this.SetData("result", result)
	return float

}
func from_sign(this *DFM, value interface{}) State {
	if !valid_char(value.(rune)) {
		return end
	}
	return this.GetState()
}

func from_number(this *DFM, value interface{}) State {
	if !unicode.IsNumber(value.(rune)) && value.(rune) != '.' && value.(rune) != '-' && value.(rune) != '+' {
		return end
	}
	if value.(rune) == '.' {
		return float
	}
	if value.(rune) == '-' {
		this.SetData("sign", float32(-1))
		return sign
	} else if value.(rune) == '+' {
		return sign
	}
	result := this.GetData("result").(float32)
	result = result*10 + float32(value.(rune)-'0')
	this.SetData("result", result)
	return number
}

func myAtoiReverse(s string) float32 {
	state_trans := map[State]StateTran{
		start:  from_start,
		sign:   from_sign,
		number: from_number,
		float:  from_float,
	}
	stats_data := map[string]interface{}{
		"sign":       float32(1),
		"result":     float32(0),
		"float_base": float32(1),
	}
	dfm := NewDFM(start, state_trans, stats_data)
	for i := len(s) - 1; i > 0; i-- {
		if dfm.GetState() == end {
			break
		}
		dfm.Step(rune(s[i]))
	}
	return dfm.GetData("result").(float32) * dfm.GetData("sign").(float32)
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
		fmt.Printf("%s 转换: %f\n", num_str, num)
	}
}
