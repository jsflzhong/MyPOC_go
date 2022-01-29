package leetcode

func reverse7(x int) int {
	tmp := 0
	for x != 0 {
		//x%10: 取最后一位数字, 即个位数. 例如123中的3.
		//在for中, 这里反复每次都取出最后一位数字放到tmp里, 然后每次把tmp后面加一位0来存它.
		tmp = tmp*10 + x%10
		//x/10: 取除了最后一位的前N位数. 例如123中的12. 即, 上面拿到了最后一位, 这里要把它从原数字中刨除出去.
		//当个位数/10时=0, 所以用这里和for的条件, 来控制从右往左的取数,直到取到第一位结束.
		x = x / 10
	}
	if tmp > 1<<31-1 || tmp < -(1<<31) {
		return 0
	}
	return tmp
}

func My_reverse(num int) int {
	tmp := 0
	for num != 0 {
		//把tmp增加一个进位,拿num的最后一位, 放入tmp新增的进位
		tmp = tmp*10 + num%10
		//刨除掉十进制num的最后一位
		num = num / 10
		//判断边界
	}
	if tmp > 1<<31-1 || tmp < -(1<<31) {
		return 0
	}
	return tmp
}
