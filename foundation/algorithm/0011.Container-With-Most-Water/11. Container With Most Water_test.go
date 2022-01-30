package leetcode

import (
	"fmt"
	"testing"
)

type Question11 struct {
	Param11
	Answer11
}

type Param11 struct {
	param []int
}

type Answer11 struct {
	answer int
}

func Test_my_maxArea(t *testing.T) {
	question11 := []Question11{
		{
			Param11{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			Answer11{49},
		},
		{
			Param11{[]int{1, 2, 3, 9, 8, 11}},
			Answer11{18},
		},
	}
	for _, v := range question11 {
		//result := maxArea(v.param)
		result := my_maxArea(v.param)
		fmt.Printf("@@@[answer]:%t, [input]:%v, [output]:%d \n",
			result == v.answer, v.param, result)
	}
}
