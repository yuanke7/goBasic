package main

import "fmt"

func main() {
	number := 0
	increment := func(amount int) {
		number += amount
	}
	increment(10)
	increment(20)
	fmt.Println(number)
}
