package main

import (
	"fmt"
	"time"
)

func main() {
	t0 := time.Now()
	time.Sleep(10 * time.Second)
	elapsed := time.Since(t0)
	fmt.Println(elapsed)
}
