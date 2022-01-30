package leetcode

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res, i := "", 0
	for num != 0 {
		for values[i] > num {
			i++
		}
		num -= values[i]
		res += symbols[i]
	}
	return res
}

// 用贪心算法, 先建模, 将 1-3999 范围内的特殊的罗马数字从大到小放在数组中，从头选择到尾，即可把整数转成罗马数字
// 具体:
//	1.迭代入参, 在贪心数组中每次找到比它小的罗马字符, 比如入参是7, 数组中比它小的是5, 从而找到对应的罗马字符V, 剩余2.
//	2.迭代剩余的入参2, 同样找到比它小的罗马字符, 是1, 即I, 剩余1.
//	3.迭代剩余的入参1, 同样找到比它小的罗马字符, 是1, 即I, 剩余0. 退出循环. 至此三个罗马字符拼接后是VII.
func my_intToRoman(inputInt int) string {
	intArray := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	stringArray := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	i, symbol := 0, ""
	for inputInt > 0 {
		//1.在贪心数组中, 找到比入参数字小的数字 和 罗马符号
		for intArray[i] > inputInt {
			i++
		}
		symbol += stringArray[i]
		//如果入参直接等于贪心数组中的数字, 则这里会为0, 则外层循环直接结束. 如果还有余数, 则继续找剩余的符号在上一行进行拼接
		inputInt -= intArray[i]
	}
	return symbol
}
