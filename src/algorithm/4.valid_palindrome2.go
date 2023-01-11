package main

// 去掉一个字符后可以是回文串吗 ？  贪心双指针
//  bcaca
//  ^    ^

// 检查字符串是不是回文串  是则返回的左右坐标相等  不是则需要判断去掉一个字符后还是不是回文
func findDifference(s string, left int, right int) (int, int) {
	for left < right {
		if s[left] != s[right] {
			return left, right
		}
		left += 1
		right -= 1
	}
	return left, right
}

func isPalindrome2(s string, left int, right int) bool {
	left, right = findDifference(s, left, right)
	return left >= right
}

//
func validPalindrome2(s string) bool {
	if s == "" {
		return false
	}
	left, right := findDifference(s, 0, len(s)-1)
	if left >= right {
		return true
	}
	return isPalindrome2(s, left+1, right) || isPalindrome2(s, left, right-1)
}

func main() {
	println(validPalindrome2("cupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupucu"))
}
