package _004_Median_of_Two_Sorted_Arrays

import "fmt"

// MedianOfTwoSortedArrays
// e.g: {1, 5, 7, 12, 20}, {2, 3, 8, 9, 10} // 1,2,3,5,7,8,9,10,12,20 = (7+8)/2 = 7
// README中的思路并不正确, 它缺失了一块逻辑.
// 即: nums1[midA-1] ≤ nums2[midB] && nums2[midB-1] ≤ nums1[midA] 这个思路还少了一部分,
// 在array1的左边nums1[midA-1] 小于 array2的右边nums2[midB] 的前提下, 还需要判断它的下一位是不是也小于array2的右边, 如果也是的话, 就需要把midA往右移
// 一直找到右边最大的符合上述条件的位置才是对的. 当然同时要判断一直到最右边的元素都是符合的的话, 再下一位就会数组越界, 所以要判断一下,
// 同时mid还不能是最后一位, 所以要判断下一位是不是尾巴了, 如果是, 就不要再右移了.
func MedianOfTwoSortedArrays(a1 []int, a2 []int) int {
	mid1, mid2, result := len(a1)>>1, len(a2)>>1, -1
	for a1[mid1-1] > a2[mid2] {
		mid1--
	}
	if a1[mid1] < a2[mid2] {
		for a1[mid1] < a2[mid2] && mid1+1 != len(a1)-1 {
			mid1++
		}
	} else {
		for a2[mid2] < a1[mid1] && mid2+1 != len(a2)-1 {
			mid2++
		}
	}
	fmt.Printf("@@@The mid1=%d, mid2=%d, mid1-1=%d, mid2-1=%d \n", mid1, mid2, a1[mid1-1], a2[mid2-1])
	if (len(a1)+len(a2))%2 == 0 {
		result = (max(a1[mid1-1], a2[mid2-1]) + min(a1[mid1], a2[mid2])) / 2
		fmt.Printf("@@@偶数, 公示: max(%d, %d) + min(%d, %d) \n", a1[mid1-1], a2[mid2-1], a1[mid1], a2[mid2])
	} else {
		result = max(a1[mid1-1], a2[mid2-1])
	}
	return result
}

func max(i1 int, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}

func min(i1 int, i2 int) int {
	if i1 > i2 {
		return i2
	}
	return i1
}
