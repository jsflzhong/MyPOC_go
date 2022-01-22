package _003_Longest_Substring_without_Repeating_Characters

// Input: "abccbabade"
// Output: 4
// The answer is "bade"
// 解法二 滑动窗口
// 重点思路:
// 1.活用: 字符串的ASC - 'a': 每个字符从0开始的特有int数字
// 2.字符串[下标]: 指定字符的asc码
func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}
	//记录某个字符的ASI码是否出现过
	var freq [256]int
	//result: 最长子串的长度
	result, left, right := 0, 0, -1

	for left < len(s) {
		//s[right+1]: 原串中, 右指针的下一位字符的ASI码
		//s[right+1]-'a': 例如下一位是c, 那么这里就是99-97=2
		//freq[s[right+1]-'a'] == 0: 下一位的ASI减a的ASC得到的数, 在freq这个数组里是否存在. 下面一行的++会把数组中这个位置的值从零设为1.
		//本行解释: 如果右指针不大于原串边界, 且右指针的下一位, 不在自己的freq数组中.
		if right+1 < len(s) && freq[s[right+1]-'a'] == 0 {
			//把原串中, 右指针的下一位字符-a后的ASC码值作为下标, 放在freq数组中, 值为1(go可以: 数组[下标]++ 的语法把数组中的某个值直接+1)
			freq[s[right+1]-'a']++
			right++
		} else {
			//把原串中, 左指针指向的字符,从freq数组中删掉, 移动窗口的左指针.
			freq[s[left]-'a']--
			left++
		}
		//记录最长的子串的长度.
		result = max(result, right-left+1)
	}
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// exercise, Input: "abccbabade"
// 返回最长子串的长度, 和其本身
func lengthOfLongestSubstring2(s string) (int, string) {
	if len(s) <= 0 {
		return -1, ""
	}
	ascArray := [256]int{}
	leftP, rightP, maxLength := 0, -1, 0
	for leftP <= len(s) && rightP+1 < len(s) {
		if ascArray[s[rightP+1]-'a'] == 0 {
			ascArray[s[rightP+1]-'a']++
			rightP++
		} else {
			ascArray[s[leftP]-'a']--
			leftP++
		}
		maxLength = max(maxLength, rightP-leftP+1)
	}
	return maxLength, s[leftP : rightP+1]
}
