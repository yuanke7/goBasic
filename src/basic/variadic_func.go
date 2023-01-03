package main

import "fmt"

// 变化参数
func average(numbers ...float64) float64 {
	total := 0.0
	for _, number := range numbers {
		//fmt.Println(index)
		total += number
	}
	return total / float64(len(numbers))
}

func main() {
	fmt.Println(average(1, 2, 3, 4, 5))
	max := 10
	//panic(fmt.Printf("The max. number is %d", max))  // Printf 不支持作为参数传递 err
	panic(fmt.Sprintf("The max. number is %d", max)) // Sprintf 支持
}
