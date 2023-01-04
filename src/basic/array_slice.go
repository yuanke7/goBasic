package main

import "fmt"

func main() {
	// 初始化数组
	var ar1 [5]int
	fmt.Println(ar1)
	// 改值
	ar1[0] = 12
	fmt.Println(ar1)
	// 由数组创建切片
	slic1 := ar1[0:3]
	fmt.Println(slic1)
	// length
	fmt.Println(len(slic1))

	// 初始化切片
	var sli []float64
	sli = append(sli, 10.012)
	fmt.Println(sli)
	sli[0] = 1.1111
	fmt.Println(sli)
}
