package main

import (
	"fmt"
	"sort"
)

//nums = [2,7,11,15], target = 9
// 哈希表解法  o(n)
func twoSum(nums []int, target int) []int {
	// 定义一个hashmap
	var numIndexMap = make(map[int]int) // 键值对都为int
	for i, num := range nums {
		if index, ok := numIndexMap[target-num]; ok {
			return []int{index, i}
		}
		numIndexMap[num] = i
	}
	return []int{}
}

// 排序双指针
type dataType [][]int

func (data dataType) Len() int {
	return len(data)
}

func (data dataType) Less(i, j int) bool {
	return data[i][1] < data[j][1] // 按二维数组中数组的第二个元素排序
}

func (data dataType) Swap(i, j int) {
	data[i], data[j] = data[j], data[i]
}

func twoSum2(nums []int, target int) []int {
	// 记录一下原始下标
	numIndex := [][]int{}
	for index, num := range nums {
		numIndex = append(numIndex, []int{index, num})
	}
	sort.Sort(dataType(numIndex))
	fmt.Println(numIndex)
	// 排序占 o(logn)
	//sort.Ints(nums)  排序数组
	// 排序二维数组
	left, right := 0, len(nums)-1
	for left < right {
		sum := numIndex[left][1] + numIndex[right][1]
		if sum == target {
			res := []int{numIndex[left][0], numIndex[right][0]}
			sort.Ints(res) // 对结果进行排序
			return res
		} else if sum < target {
			left += 1
		} else if sum > target {
			right -= 1
		}
	}
	return []int{}
}

func main() {
	nums := []int{11, 15, 2, 7, 6, 8}
	//fmt.Println(twoSum(nums, 9))
	fmt.Println(twoSum2(nums, 17))
}
