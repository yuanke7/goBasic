package main

import (
	"fmt"
	"strings"
	"unicode"
)

func validPalindrome(s string) bool {
	// 检测合法字符
	s = strings.TrimSpace(s)
	newS := ""
	for i := range s {
		if unicode.IsLetter(rune(s[i])) || unicode.IsNumber(rune(s[i])) {
			newS += string(unicode.ToLower(rune(s[i])))
		}
	}
	fmt.Println(newS)
	if newS == "" {
		return true
	}

	left, right := 0, len(newS)-1
	for left <= right {
		leftVal := string(newS[left])
		rightVal := string(newS[right])
		if leftVal != rightVal {
			return false
		}
		left += 1
		right -= 1
	}
	return true
}

func main() {
	s := "`l;`` 1o1 ??;l`"
	print(validPalindrome(s))
	//fmt.Println(unicode.IsPunct(""))
}
