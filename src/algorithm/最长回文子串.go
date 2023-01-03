package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	s := "abcba"
	isPalindrome(s)
	//longestPalindrome("dsadsa")
}

func isPalindrome(s string) bool {
	leftIndex := math.Floor(float64(len(s)/2) - 1) // 地板除
	rightIndex := math.Floor(float64(len(s)/2) + 1)
	fmt.Println(leftIndex, rightIndex, reflect.TypeOf(leftIndex), reflect.TypeOf(rightIndex)) // 查看变量类型
	for {
		//if s[leftIndex] ==
		break
	}
	return true
}
func longestPalindrome(s string) string {
	tmp := ""
	for i := 0; i < len(s); i++ {
		fmt.Println(string(s[i]))
		tmp = tmp + string(s[i])

	}
	return ""
}
