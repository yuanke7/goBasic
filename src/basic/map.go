package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func str2int(s int) string {
	i := strconv.Itoa(s)
	//if err != nil {
	//	panic("String " + s + "is not a number!")
	//}
	return i
}

func main() {
	//a := make(map[string]int)
	a := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	a["name"] = 857
	fmt.Println(a)
	//println(a["name"])
	//println(a)        // 打印地址
	//delete(a, "name") // remove key
	//println(a["name"])
	//println(a)
	//num, ok := a["name"]
	//fmt.Println(num, ok)
	//a["name"] = 9527
	num, ok := a["name"]
	if ok { // 判断键值对存在
		fmt.Println(num, ok)
	}
	res := str2int(123)
	fmt.Println(reflect.TypeOf(res))
}
