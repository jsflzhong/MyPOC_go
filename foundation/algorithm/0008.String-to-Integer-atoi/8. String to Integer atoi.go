package leetcode

//别人的答案, 自己的分析注释
func myAtoi(s string) int {
	//maxInt: int最大边界值
	//signAllowed: 是否允许正负号. (只在开头第一个字符之前的正负号才被允许)
	//whitespaceAllowed: 是否允许空格. (只在开头第一个字符之前的空格才被允许)
	//sign: 正负号. 以1或-1表示. 最后直接乘以结果, 就可以了.
	//digits: 用来存放迭代字符串之后取出的数字.
	maxInt, signAllowed, whitespaceAllowed, sign, digits := int64(2<<30), true, true, 1, []int{}

	//1.迭代字符串, 放入数组.
	//在ASCII中, 字符 '0到9' 对应着十进制数字为: '48到57'. 例如字符'4'对应的十进制就是:52. ###注意, 差值是'48'.
	for _, c := range s {
		if c == ' ' && whitespaceAllowed {
			continue
		}
		if signAllowed {
			if c == '+' {
				signAllowed = false
				//开头正负号之后的空格, 也不允许读取了
				whitespaceAllowed = false
				continue
			} else if c == '-' {
				sign = -1
				signAllowed = false
				whitespaceAllowed = false
				continue
			}
		}
		//这样检查是否不是数字? --是的, 用字符和'0'对比, 实际上对比的是ASCII吗, 字符'0'的十进制是48, 字符'9'的十进制是57. 注意字符之间是可以直接用大于小于号来对比的
		if c < '0' || c > '9' {
			break
		}
		whitespaceAllowed, signAllowed = false, false
		//int(c-48) = ASCII中字符对应的十进制, 与其实际的十进制, 差了48. 例如字符'4'的ASCII十进制是52, 52-48=4(这是int了)
		digits = append(digits, int(c-48))
	}
	var num, place int64
	place, num = 1, 0

	//2.处理数字开头的零
	lastLeading0Index := -1
	for i, d := range digits {
		if d == 0 {
			lastLeading0Index = i
		} else {
			break
		}
	}
	if lastLeading0Index > -1 {
		digits = digits[lastLeading0Index+1:]
	}

	//3.确定正负两端最大的边界
	//由于需求是: range [-2^31, 2^31 - 1]
	//所以, 当数字式正数时, 需要最后-1; 当是负数时则不用这样处理.
	var rtnMax int64
	if sign > 0 {
		rtnMax = maxInt - 1
	} else {
		rtnMax = maxInt
	}

	//4.int数组转int.
	// 倒序迭代上面组合的数组, 加和成最终的int数字.
	digitsCount := len(digits)
	for i := digitsCount - 1; i >= 0; i-- {
		//int64(digits[i])= 把int转int64后参与计算
		//迭代int数组, 十进位
		num += int64(digits[i]) * place
		place *= 10
		//判断是否超出range [-2^31, 2^31 - 1] , 即: [-2147483648, 2147483648]
		if digitsCount-i > 10 || num > rtnMax {
			//int64(sign) = sign代表正负号, 是1或-1两种, 转变成int64后参与计算, 所以返回的是正的还是负的2^30, 是由这里控制的.
			return int(int64(sign) * rtnMax)
		}
	}

	//5.乘上正负号,结束.
	num *= int64(sign)
	return int(num)
}

// My_Atoi
// 自己的答案
func My_Atoi(s string) int {
	//startSign: 是开头的正负号么, 是的话就允许读入该正负号.
	//startSpace: 是开头的空格么, 是的话就允许读入该空格.
	//sign: 正负号, 用1或-1标识, 最后乘上结果.
	startSign, startSpace, sign, intArray := true, true, 1, []int{}

	//1.迭代字符串, 把其中的数字放入int数组.
	//需要处理字符串开头的空格, 正负号等.
	for _, v := range s {
		//Read in and ignore any leading whitespace.
		if startSpace == true && ' ' == v {
			continue
		}
		//只有字符最左边开头的第一个正负号允许被读取
		if startSign == true {
			if '-' == v {
				//这里已经读到一个负号, 如果后面如果还有, 就不允许再读取了.
				startSign = false
				//开头正负号之后的空格, 也不允许读取了
				startSpace = false
				sign = -1
				continue
			}
			if '+' == v {
				startSign = false
				startSpace = false
				continue
			}
		}
		//处理防止中间部分出现正负号或空格的情况
		startSign, startSpace = false, false
		//遇到非数字部分就停止
		if v < '0' || v > '9' {
			break
		}
		intArray = append(intArray, int(v-48))
	}

	//2.处理int数组"开头的"零
	//Convert these digits into an integer (i.e. "123" -> 123, "0032" -> 32). If no digits were read, then the integer is 0. Change the sign as necessary (from step 2).
	indexOfZero := -1
	for i, v := range intArray {
		if 0 == v {
			indexOfZero = i
			continue
		} else {
			//仅处理开头的零
			break
		}
	}
	if indexOfZero != -1 {
		intArray = intArray[indexOfZero+1:]
	}

	//3.定义正负两端的边界
	//If the integer is out of the 32-bit signed integer range [-231, 231 - 1], then clamp the integer so that it remains in the range. Specifically, integers less than 231 should be clamped to 231, and integers greater than 231 - 1 should be clamped to 231 - 1
	boundary := 2 << 30
	if sign == 1 {
		boundary -= 1
	} else {
		boundary *= sign
	}

	//4.int[] 转 int
	resultNum := 0
	for _, v := range intArray {
		resultNum = resultNum*10 + v
	}

	//5.处理正负
	resultNum *= sign

	//6.判断是否超出边界
	if (sign == 1 && resultNum > boundary) || (sign == -1 && resultNum < boundary) {
		return boundary
	}

	return resultNum
}
