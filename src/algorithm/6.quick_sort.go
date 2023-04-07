package main

import "fmt"

// 平均时间复杂度 o(nlogn)  根据pivot 的选择， 或快或慢
// 思路：选标杆 原地算法（不需要额外空间） 大于标杆放右边  小于放左边  先局部有序再整体有序
func main() {
	// 待排序的数组
	arr := []int{5, 2, 4, 9, 7}

	// 对整个数组进行排序
	quickSortA(arr, 0, len(arr)-1)

	// 打印排序后的数组
	fmt.Println(arr)
}

func quickSort(arr []int, left int, right int) {
	if left > right {
		return
	}
	pivot := arr[right]
	partitionIndex := left
	for i := left; i < right; i++ {
		if arr[i] < pivot {
			arr[i], arr[partitionIndex] = arr[partitionIndex], arr[i]
			partitionIndex++
		}
	}
	arr[partitionIndex], arr[right] = arr[right], arr[partitionIndex]

	quickSort(arr, left, partitionIndex-1)
	quickSort(arr, partitionIndex+1, right)
}

func quickSortA(arr []int, left int, right int) {
	if left > right {
		return
	}
	pivot := arr[right]
	partitionIndex := left
	for i := left; i < right; i++ {
		if arr[i] < pivot {
			arr[i], arr[partitionIndex] = arr[partitionIndex], arr[i]
			partitionIndex++
		}
	}
	arr[partitionIndex], arr[right] = arr[right], arr[partitionIndex]

	quickSort(arr, left, partitionIndex-1)
	quickSort(arr, partitionIndex+1, right)
}
