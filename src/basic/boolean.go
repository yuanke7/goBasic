package main

import "fmt"

func main() {
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!false)
	x := 1
	if x != 0 {
		fmt.Println("Yes")
	}
	var y []string
	if len(y) != 0 {
		fmt.Println("this won't be printed")
	}
}
