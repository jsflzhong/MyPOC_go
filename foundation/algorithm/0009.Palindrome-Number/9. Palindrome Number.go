package leetcode

import "strconv"

// 解法一
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	arr := make([]int, 0, 32)
	for x > 0 {
		arr = append(arr, x%10)
		x = x / 10
	}
	sz := len(arr)
	for i, j := 0, sz-1; i <= j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
}

// 解法二 数字转字符串
func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	s := strconv.Itoa(x)
	length := len(s)
	for i := 0; i <= length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}

func my_isPalindrome_string(i int) bool {
	//leftP/rightP: 用来向右左扩散的指针
	//leftL/rightL: 用来记录最大扩散结果的指针
	leftP, rightP, leftL, rightL := 0, -1, 0, 0
	s := strconv.Itoa(i)
	length := len(s)
	for leftP < length {
		//1.处理连续相同的字符,扩散右指针. 例如111
		for rightP+1 < length && s[rightP+1] == s[leftP] {
			rightP++
		}
		//2.处理左右成对相同的字符, 左右一起扩散. 例如:121
		if rightP+1 < length && leftP-1 >= 0 && s[rightP+1] == s[leftP-1] {
			leftP--
			rightP++
		}
		//3.判断并保留扩散的结果.
		if rightP-leftP > rightL-leftL {
			leftL, rightL = leftP, rightP
		}
		//4.重置(向右移)扩散的中轴
		leftP = (leftP+rightP)/2 + 1
		rightP = leftP
	}
	if leftL == 0 && rightL == length-1 {
		return true
	}
	return false
}

func my_isPalindrome_string2(i int) bool {
	s := strconv.Itoa(i)
	length := len(s)
	str := ""
	x := length - 1
	for x >= 0 {
		str = str + string(s[x])
		x--
	}
	if str == s {
		return true
	}
	return false
}

func my_isPalindrome_int(i int) bool {
	if i < 0 {
		return false
	}
	if i == 0 {
		return true
	}
	if i%10 == 0 {
		return false
	}
	intArray := make([]int, 0, 32)
	x := i
	for x > 0 {
		lastNum := x % 10
		intArray = append(intArray, lastNum)
		x = x / 10
	}
	intValue := 0
	for _, v := range intArray {
		intValue = intValue*10 + v
	}
	if intValue == i {
		return true
	}
	return false
}
