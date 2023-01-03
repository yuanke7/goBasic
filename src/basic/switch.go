package main

import "fmt"

func main() {
	var num string
	_, err := fmt.Scanln(&num)
	if err != nil {
		return
	}
	switch num {
	case "1":
		fmt.Printf("SHIT: %s", num)
	case "2":
		fmt.Printf("SHIT: %s", num)
	case "10":
		fmt.Printf("Yes: %s", num)
	default:
		fmt.Println(nil)
	}

	// alternative solution
	_, err2 := fmt.Scanln(&num)
	if err2 != nil {
		return
	}
	num_map := map[string]string{
		"1":  "SHIT: " + num,
		"2":  "SHIT: " + num,
		"10": "Yes: " + num,
	}
	val, ex := num_map[num]
	if ex {
		fmt.Println(val)
	} else {
		fmt.Println(nil)
	}
}
