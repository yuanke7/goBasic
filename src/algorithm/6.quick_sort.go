package main

import "fmt"

func main() {
	// 待排序的数组
	arr := []int{5, 2, 4, 9, 7}

	// 对整个数组进行排序
	quickSort(arr, 0, len(arr)-1)

	// 打印排序后的数组
	fmt.Println(arr)
}

// 再写一遍
func quickSort(arr []int, left int, right int) {
	if left > right {
		return
	}
	pivot := arr[right]
	partitionIndex := left
	for i := left; i < right; i++ {
		if arr[i] < pivot {
			arr[partitionIndex], arr[i] = arr[i], arr[partitionIndex]
			partitionIndex++
		}
	}
	arr[partitionIndex], arr[right] = arr[right], arr[partitionIndex]
	quickSort(arr, left, partitionIndex-1)
	quickSort(arr, partitionIndex+1, right)
}
