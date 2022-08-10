package main

import (
	"fmt"
)

func main() {
	var num float32
	fmt.Scan(&num)

	left := int(num)
	right := int(num*10) % 10
	result := left
	if right >= 5 {
		result = left + 1
	}
	fmt.Println(result)
}
