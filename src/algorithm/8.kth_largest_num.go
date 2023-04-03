package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := []int{2, 4, 2, 5, 6, 7, 9, 12}
	res := quickSelectA(arr, 0)
	fmt.Println(res)
}

// 第K个最大的数 快速选择算法 时间复杂度 o(n)
func quickSelect(arr []int, k int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	// 随机选取一个标杆元素
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(arr))
	pivot := arr[randomIndex]

	// 把数组分为比标杆大的和比标杆小的  遍历一遍，子函数也只是遍历原数组的一部分，差不多就是o(n)
	left, right := []int{}, []int{}
	for _, item := range arr {
		if item < pivot {
			left = append(left, item)
		} else if item > pivot {
			right = append(right, item)
		}
	}
	fmt.Println("left", left)
	fmt.Println("right", right)

	// 递归处理两边的
	right_len := len(right)
	if k > right_len { // 大于大数大数组的长度，说明要找的数在左边
		return quickSelect(left, k-right_len-1) // 减去标杆和右边数组的长度
	} else if k == right_len {
		return pivot
	} else { // 小于大数大数组的长度，说明要找的数在大数数组中， 在右边接着找
		return quickSelect(right, k)
	}
}

func quickSelectA(arr []int, k int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	// 随机选取一个标杆
	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn(len(arr))
	pivot := arr[randIndex]
	left, right := []int{}, []int{}
	for _, item := range arr {
		if item < pivot {
			left = append(left, item)
		} else if item > pivot {
			right = append(right, item)
		}
	}

	// 找出第K大的元素
	rightLen := len(right)
	if k == rightLen {
		return pivot
	} else if k > rightLen {
		return quickSelectA(left, k-1-rightLen)
	} else {
		return quickSelectA(right, k)
	}
}
