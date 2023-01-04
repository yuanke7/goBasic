package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// 在一个string中找另一个子串 o(n^2)
func findSubString(source string, target string) []int {
	start, end, index := "", "", 0
	tmpIndex := 0
	for ; start == end && index < len(source)-1; index++ {
		tmpIndex = index
		for i := range target {
			if source[tmpIndex] == target[i] {
				if start == "" {
					start = strconv.Itoa(tmpIndex)
				} else if len(target)-1 == i {
					end = strconv.Itoa(tmpIndex)
					break
				}
			} else {
				start, end = "", ""
				break
			}
			tmpIndex += 1
		}
	}
	a, _ := strconv.Atoi(start)
	b, _ := strconv.Atoi(end)
	res := []int{a, b}
	fmt.Println(reflect.TypeOf(res), res)
	return res
}

// Rabin-Karp 算法实现  思想：把子串与数字对应， 快速判断子串是否存在于hashmap
func rabinKarp() {

}

func main() {
	findSubString("1542", "4")
}
