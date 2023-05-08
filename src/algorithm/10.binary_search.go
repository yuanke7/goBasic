package main

func main() {
	arr := []int{1, 2, 8, 9, 10, 10, 12}
	//println(binarySearchRecursion(arr, 0, len(arr)-1, 1))
	//println(binarySearchNormal(arr, 12))
	println(binaryModel(arr, 9))
}

// 递归解法  o(logN)  是一种减治的思想，而不是分治
func binarySearchRecursion(arr []int, start int, end int, target int) int {
	// 跳出条件 一定要注意 否则死循环
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

// 循环解法
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

// 二分法模板 todo 重点记忆
func binaryModel(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	start, end := 0, len(nums)-1

	for start+1 < end { // 双指针相邻时就退出
		mid := (start + end) / 2

		if nums[mid] < target {
			start = mid
		} else if nums[mid] == target {
			end = mid
		} else {
			end = mid
		}
	}

	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	return 1
}
