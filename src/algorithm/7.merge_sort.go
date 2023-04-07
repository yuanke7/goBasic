package main

import "fmt"

// 时间复杂度 o(nlogn) 问题规模每次减半
// 思路：直接取数组中部的元素作为标杆， 左右数组

func main() {
	arr := []int{5, 2, 4, 6, 1, 3}
	sorted := mergeSortA(arr)
	fmt.Println(sorted)
	//lst := make([][]int, 0, 10)
	//lst = append(lst, arr[2:])
	//fmt.Println(lst)
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

func merge(left, right []int) []int {
	i, j := 0, 0
	result := make([]int, 0, len(left)+len(right))
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func mergeSortA(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSortA(arr[:mid])
	right := mergeSortA(arr[mid:])
	return mergeA(left, right)
}

func mergeA(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
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
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
