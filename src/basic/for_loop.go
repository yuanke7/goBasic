package main

import "fmt"

func main() {
	names := []string{"a", "b", "c"}
	for index, name := range names { // 相当于Python enumerate
		fmt.Printf("index:%d, name:%s\n", index, name)
	}

	for a := 1; a < 1000; a++ { // 相当于while
		//fmt.Println(1)
	}
}
