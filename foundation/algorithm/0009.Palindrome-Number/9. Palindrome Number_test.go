package leetcode

import (
	"fmt"
	"testing"
)

type Question9 struct {
	Param9
	Answer9
}

type Param9 struct {
	param int
}

type Answer9 struct {
	answer bool
}

func Test_PalindromeNumber(t *testing.T) {
	questionArray := []Question9{
		{
			Param9{121},
			Answer9{true},
		},
		{
			Param9{-121},
			Answer9{false},
		},
		{
			Param9{10},
			Answer9{false},
		},
		{
			Param9{111},
			Answer9{true},
		},
	}

	for _, v := range questionArray {
		//result := my_isPalindrome_string(v.param)
		//result := my_isPalindrome_string2(v.param)
		result := my_isPalindrome_int(v.param)
		fmt.Printf("@@@[result]: %t,   [input]:%v,  [output]:%v \n", result == v.answer, v.param, result)
	}
}
