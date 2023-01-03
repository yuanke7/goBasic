package main

import (
	"fmt"
	"math/rand"
)

func quickSort(values []int) []int {
	values[0] = 100000
	return values
}

var global_var = 1 // 通常用来作为全局变量

func main() {
	// 指定容量则为数组 不可变
	fmt.Println(global_var)
	a := [10]int{} // 海象符通常作为局部变量
	a[1] = 321321
	fmt.Println(a)
	rand.Seed(1)
	for i := 0; i < 10; i++ {
		a[i] = rand.Intn(10000)
	}
	as := a[:] // go中的切片是一个引用
	//fmt.Println(a, as)
	quickSort(as)
	//fmt.Println(&a[0], &as[0]) //切片与数组的第一个元素地址相同
	//fmt.Println(as, a, len(a), len(as))
	// 切片更加易用
	b := []int{}
	b = append(b, 2000)
	fmt.Println(b)
	b[0] = 100
	fmt.Println(b)
}
