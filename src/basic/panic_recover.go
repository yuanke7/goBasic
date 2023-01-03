package main

import "fmt"

func try() {
	fmt.Println("Recover ...")
	p := recover()
	if p != nil {
		fmt.Printf("Panic: %s", p)
	}
	fmt.Println("End.")
}

func main() {
	// defer 语句延时逆序执行
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover 捕获到了异常", err)
		}
	}()
	panic("Shit!")
}
