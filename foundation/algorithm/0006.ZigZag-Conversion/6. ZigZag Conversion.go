package leetcode

//Skip
func convert(s string, numRows int) string {
	matrix, down, up := make([][]byte, numRows, numRows), 0, numRows-2
	for i := 0; i != len(s); {
		if down != numRows {
			matrix[down] = append(matrix[down], byte(s[i]))
			down++
			i++
		} else if up > 0 {
			matrix[up] = append(matrix[up], byte(s[i]))
			up--
			i++
		} else {
			up = numRows - 2
			down = 0
		}
	}
	solution := make([]byte, 0, len(s))
	for _, row := range matrix {
		for _, item := range row {
			solution = append(solution, item)
		}
	}
	return string(solution)
}

// My_convert
// Z型(齿轮型 zigzag pattern )输出
//比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
//P   A   H   N
//A P L S I I G
//Y   I   R
//规律分析:
//在num=3时,
//第一行的下标为0,空,4,空,8,... 即 第一次是0(行数-1), 后面是+n+1, 每2数之间隔空
//第二行的下标为1,3,5,7,9,11,... 即 第一次是1(行数-1), 从第二次开始,是+n-1
//第三行的下标为2,6,10,14,... 即 第一次是2(行数-1), 从第二次开始,是+n+1, 每2数之间隔空
//
//如果num=4
//P     I    N
//A   L S  I G
//Y A   H R
//P     I
//规律:
//1.每行的长度, 即列的个数: 固定为7?
//2.每行空位的个数: n-2
//3.当num改变时,每行的规律是不同的: 1行:
func My_convert(s string, num int) string {
	return ""
}
