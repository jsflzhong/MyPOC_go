package _005_Longest_Palindromic_Substring

import (
	"fmt"
	"testing"
)

func Test_LongestPalindromicSubstring(t *testing.T) {
	//s := "abbcaddacert" //caddac
	s := "abcccbc"
	result := longestPalindrome1(s)
	fmt.Printf("@@@Test_LongestPalindromicSubstring Result is :%s \n", result)
}

func Test_My_longestPalindrome1(t *testing.T) {
	s := "abbcaddacert" //caddac
	//s := "abcccbc"  //bcccb
	result := My_longestPalindrome1(s)
	fmt.Printf("@@@Test_My_longestPalindrome1 Result is :%s \n", result)
}
