package main

import (
	"fmt"
	"reflect"
)

// 最长回文子串
func isPalindrome(s string) bool {
	if s == "" {
		return false
	}
	leftIndex, rightIndex := 0, len(s)-1
	for {
		if s[leftIndex] != s[rightIndex] {
			return false
		} else {
			leftIndex += 1
			rightIndex -= 1
		}
		if leftIndex >= rightIndex {
			return true
		}
	}

}

// 暴力穷举o(n^3)
func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	var maxLen = 0
	var longestSubStr string
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s)+1; j++ {
			cur := s[i:j]
			//fmt.Println(i, j, cur)
			if isPalindrome(cur) && len(cur) > maxLen {
				longestSubStr = cur
				maxLen = len(cur)
			}
		}
	}
	return longestSubStr
}

// 中心扩散法 o(n^2)
func longestPalindromeMidSpray(s string) string {
	if len(s) <= 1 {
		return s
	}
	maxLen := 0
	maxStr := ""
	sLen := len(s)
	var spray = func(left int, right int) {
		var tmpStr = ""
		for {
			if left < 0 || right > sLen-1 {
				break
			}
			if s[left] == s[right] {
				if left == right {
					tmpStr = string(s[right])
				} else {
					tmpStr = s[left : right+1]
				}
				if len(tmpStr) > maxLen {
					maxStr = tmpStr
					maxLen = len(tmpStr)
				}
				left -= 1
				right += 1
			} else {
				break
			}
		}
	}

	for i := 0; i < sLen; i++ {
		fmt.Println(string(s[i]))
		// 以 i 为中心扩散
		var tmpLeft = i
		var tmpRight = i
		spray(tmpLeft, tmpRight)
		// 以i+0.5为中心扩散 i i+1 为初始左右游标
		//tmpLeft = i
		tmpRight = i + 1
		spray(tmpLeft, tmpRight)
	}

	return maxStr
}

// 动态规划dp o(n^2)
func longestPalindromeDp(s string) string {
	maxStr := ""
	n := len(s)
	// 初始化二维数组  在【区间型】动态规划中，以二维数组的坐标值作为 区域（string or slice） 的状态
	isPalindrome := make([][]bool, n)
	for i := range isPalindrome {
		isPalindrome[i] = make([]bool, n)
		isPalindrome[i][i] = true // 单个元素本身是一个回文串
	}
	for i := range isPalindrome {
		if i == 0 {
			continue
		}
		isPalindrome[i][i-1] = true
	}

	fmt.Println(isPalindrome)
	return maxStr
}

func main() {
	s := "abbcccda"

	//println(isPalindrome(s))
	//var res = longestPalindrome(s)
	//var res = longestPalindromeMidSpray(s)
	var res = longestPalindromeDp(s)
	fmt.Println(res, reflect.TypeOf(res))
}
