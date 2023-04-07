package main

// 递归三要素
// 1. 定义：函数要接受什么样的参数  返回什么样的值，代表什么意思
// 2. 递归的拆解： 通过处理把参数规模变小
// 3. 递归的出口： 适时结束递归

func main() {
	print(fib(5))
}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
