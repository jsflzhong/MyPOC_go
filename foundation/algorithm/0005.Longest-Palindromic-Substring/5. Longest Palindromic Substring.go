package _005_Longest_Palindromic_Substring

// LongestPalindromicSubstring
// ###回文字符串的特点肯定是: 从左往右看, 和从右往左看, 是一样的. 但不表示一定是两两数字对称的, 看下面的例2.

// 例1: "abbcaddacert" =  caddac
// 例2: ""abcccbc" = bcccb 这种不是对称的也算, 因为去和回都是一样的.

// 解法三，滑动窗口
// 这个写法其实就是中心扩散法变了一个写法。中心扩散是依次枚举每一个轴心。
// 滑动窗口的方法稍微优化了一点，有些轴心两边字符不相等，下次就不会枚举这些不可能形成回文子串的轴心了。
// 不过这点优化并没有优化时间复杂度，时间复杂度 O(n^2)，空间复杂度 O(1)。
func longestPalindrome1(s string) string {
	if len(s) == 0 {
		return ""
	}
	//pr,pl用来保留最长的回文串的两个端点.
	//left, right用来保留从轴心往两端扩散的指针
	left, right, pl, pr := 0, -1, 0, 0
	for left < len(s) {
		// 移动到相同字母的最右边（如果有相同字母）
		// 例: "abbcaddacert" =  caddac
		// 例: "abcccbc" = bcccb 这种不是对称的也算, 因为去和回都是一样的.
		// left和right在下面设为了一样的值. 所这里是left和它右边的值不停的对比, 如果一样, 就把右指针右移.
		for right+1 < len(s) && s[left] == s[right+1] {
			right++
		}
		// 定好回文的边界, 然后不停对比left左边和right右边的两个值, 因为回文字符串的特点肯定是左右俩值是一样的, 例如 abc||cba
		for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
			left--
			right++
		}
		// 与上次保存的最长的回文串的指针对比, pr,pl用来保留最长的回文串的两个端点.
		if right-left > pr-pl {
			pl, pr = left, right
		}
		// 重置到下一次寻找回文的中心. 规则: 比现在扩散后的中心点再往右一个点, 就可以了. 就是本次扩散的中心点的下一个位置了. 避免了死循环.
		// 例: "abcccbc", 第三次循环后到第一个c, 然后right一直++到第三个c, 然后上面的最后一个for循环会扩散到left=1, right=5, 也就是两个b.
		//	 然后, 这里取中心3,就是本次回文串的中心, 再+1, 就从下一个开始对比了.
		left = (left+right)/2 + 1
		right = left
	}
	return s[pl : pr+1]
}

func My_longestPalindrome1(s string) string {
	if len(s) == 0 {
		return ""
	}
	leftP, rightP, leftL, rightL := 0, -1, 0, 0
	for leftP < len(s) {
		// 通过右移右指针, 找到连续重复的字符. 例如"bcccb". 如果没有, 在第一轮时就让right和left合并在一个位置上.
		for rightP+1 < len(s) && s[rightP+1] == s[leftP] {
			//因为这里要右移右指针, 所以上面要加上右指针不越界的条件
			rightP++
		}
		// 定义边界, 并开始扩散
		for leftP-1 >= 0 && rightP+1 < len(s) && s[leftP-1] == s[rightP+1] {
			leftP--
			rightP++
		}
		//上面已经找到回文串,与旧的回文串长度做对比, 保留更长的那一个.
		if rightP-leftP > rightL-leftL {
			leftL = leftP
			rightL = rightP
		}
		//重新确定下一轮的扩散中心点
		//由于这里的leftP会在每一轮的这里不停的往右移, 所以最外层的for循环的停止条件就可以确定为 if leftP < len(s) 了.
		leftP = (leftP+rightP)/2 + 1
		rightP = leftP
	}
	return s[leftL : rightL+1]
}
