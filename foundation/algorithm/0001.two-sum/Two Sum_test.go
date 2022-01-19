package _001_two_sum

import (
	"fmt"
	"testing"
)

func Test_twoSum(t *testing.T) {
	array := []int{1, 3, 5, 7, 11}
	target := 12
	result := twoSum(array, target)
	fmt.Printf("@@@Test_twoSum, [intput] is %v,   [target] is %d,  [output] is %v", array, target, result)
}
