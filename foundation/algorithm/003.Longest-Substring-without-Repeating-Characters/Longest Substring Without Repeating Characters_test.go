package _03_Longest_Substring_without_Repeating_Characters

import (
	"fmt"
	"testing"
)

func Test_lengthOfLongestSubstring1(t *testing.T) {
	s := "abccbabade"
	result := lengthOfLongestSubstring1(s)
	fmt.Printf("@@@The result is : %d \n", result)
}

func Test_lengthOfLongestSubstring2(t *testing.T) {
	//s := "abccbabade"  //bade
	s := "abccbabadeba" //deba 窗口左指针会把左边重复的字符去除掉, 所以ba会被去掉, 留下后面的deba
	result, resultString := lengthOfLongestSubstring2(s)
	fmt.Printf("@@@The maxLength is : %d, the maxSubString is: %s \n", result, resultString)
}
