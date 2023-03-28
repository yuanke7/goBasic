package main

import "fmt"

// 时间复杂度 o(nlogn) 问题规模每次减半

func main() {
	arr := []int{5, 2, 4, 6, 1, 3}
	sorted := mergeSort(arr)
	fmt.Println(sorted)
	//lst := make([][]int, 0, 10)
	//lst = append(lst, arr[2:])
	//fmt.Println(lst)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right)) // 初始化一个长度为0 容量为左右左右数组和的切片， 超出容量后会自动分配二倍的容量
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...) // ... 表示将切片解包
	result = append(result, right[j:]...)
	return result
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}
