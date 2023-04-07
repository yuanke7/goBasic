package main

func main() {
	arr := []int{1, 2, 8, 9, 10, 12}
	println(binarySearchRecursion(arr, 0, len(arr)-1, 1))
	println(binarySearchNormal(arr, 12))
}

func binarySearchRecursion(arr []int, start int, end int, target int) int {
	// 跳出条件
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if arr[mid] == target {
		return mid
	}

	// 递归处理两边
	if arr[mid] < target {
		return binarySearchRecursion(arr, mid+1, end, target)
	}
	return binarySearchRecursion(arr, start, mid-1, target)
}

func binarySearchNormal(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}

	start, end := 0, len(arr)-1
	for start < end {
		mid := (start + end) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}

	}
	if arr[start] == target {
		return start
	}

	return -1
}
