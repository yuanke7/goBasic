package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func printFunc(content string) {
	defer wg.Done() // 协程结束就-1
	fmt.Println("PrintFunc is running ...")
	time.Sleep(1 * time.Second)
	fmt.Println(content)
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("%d ", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("%c ", i)
	}
}

func main() {
	//go printFunc("SHIT") // 这行开跑后立刻跳到下一行，不等待本行执行完毕
	////time.Sleep(1 * time.Second)
	//go numbers()
	//go alphabets()
	//time.Sleep(10 * time.Second)
	//fmt.Println("main function") // 只要执行到这里就完了，不管 gorutine 是否完毕
	// -------------------------------
	// 不用sleep 实现等待所有goroutine完成
	urls := []string{
		"https://www.peterbe.com",
		"https://python.org",
		"https://golang.org",
	}
	for _, url := range urls {
		wg.Add(1) // 开一个协程就在组内+1
		go printFunc(url)
	}
	// Wait for the goroutines to finish
	wg.Wait()
}
