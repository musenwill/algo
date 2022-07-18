package main

import (
	"fmt"
)

func main() {
	var num int
	_, _ = fmt.Scan(&num)

	maxFactor := 2

	for num > 1 {
		i := maxFactor
		for i <= num {
			if num%i == 0 {
				num /= i
				maxFactor = i
				fmt.Printf("%d ", i)
				break
			}
			if i&0x01 == 0x01 {
				i += 2
			} else {
				i++
			}
		}
	}
}
