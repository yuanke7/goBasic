package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

// 普通参数在内部的更改不影响传入参数的值，也就是说传入的是原本参数的copy
func GetDistance(a Point, b Point) float64 {
	a.x += 1
	a.y += 1
	b.x += 2
	b.y += 2
	fmt.Println(&a.x, &a.y, &b.x, &b.y) // 0x1400001e0a0 0x1400001e0a8 0x1400001e0b0 0x1400001e0b8 因为传入的copy 所以地址不同
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}

// 使用指针参数，函数内部改变会影响传入参数的值
func GetDistanceBetter(a *Point, b *Point) float64 {
	//a.x += 1
	//a.y += 1
	//b.x += 2
	//b.y += 2
	fmt.Println(&a.x, &a.y, &b.x, &b.y) // 0x1400001e070 0x1400001e078 0x1400001e080 0x1400001e088 参数地址与传入的相同
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}

// struct method 需要改变结构体值时需要使用指针
func (this *Point) structMethodGetDistance(a Point, b Point) {
	fmt.Println(a, b, this.x, this.y)
	a.x++
	a.y++
	b.x++
	b.y++
	fmt.Println(a, b)
}

func main() {
	p1 := Point{0, 0}
	p2 := Point{0, 0}
	fmt.Println(&p1.x, &p1.y, &p2.x, &p2.y) // 0x1400001e070 0x1400001e078 0x1400001e080 0x1400001e088 参数地址
	//res1 := GetDistance(p1, p2)
	//res2 := GetDistanceBetter(&p1, &p2)
	p1.structMethodGetDistance(p1, p2)
	//fmt.Println(res1, res2)
	//fmt.Println(p1, p2)
}
