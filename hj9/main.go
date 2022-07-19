package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)

	digits := make([]int, 10)
	result := 0

	for num > 0 {
		d := num % 10
		num /= 10

		if digits[d] == 0 {
			result *= 10
			result += d
		}

		digits[d] = 1
	}

	fmt.Println(result)
}
