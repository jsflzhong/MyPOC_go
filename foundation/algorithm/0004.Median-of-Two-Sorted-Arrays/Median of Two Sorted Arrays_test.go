package _004_Median_of_Two_Sorted_Arrays

import (
	"fmt"
	"testing"
)

// 1, 2, 3, 5, 6, 7, 8, 9, 12 = 6
func Test_MedianOfTwoSortedArrays(t *testing.T) {
	//a2 := []int{1, 5, 7, 12}
	//a1 := []int{2, 3, 6, 8, 9} //1, 2, 3, 5, 6, 7, 8, 9, 12 = 6

	//a1 := []int{1, 5, 7, 12, 20}
	//a2 := []int{2, 3, 8, 9, 10} // 1,2,3,5,7,8,9,10,12,20 = (7+8)/2 = 7

	//a1 := []int{1, 5, 7, 12, 20} //测试a1直到最右边的mid1都比a2的mid2小的情况, 处理越界和不能是最后一位的情况.
	//a2 := []int{2, 3, 21, 23, 25} // 1,2,3,5,7,12,20,21,23,25 = (7+12)/2 = 9

	a1 := []int{1, 5, 7, 12, 19, 20} //处理越界+奇数的情况. 上面是偶数
	a2 := []int{2, 3, 21, 23, 25}    // 1,2,3,5,7,12,19,20,21,23,25 = 12

	result := MedianOfTwoSortedArrays(a1, a2)
	fmt.Printf("@@@The result is: %d", result)
}
