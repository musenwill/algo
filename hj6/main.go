package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	_, _ = fmt.Scan(&num)
	sqrt := int(math.Sqrt(float64(num)))

	for i := 2; i <= sqrt; {
		for num%i == 0 {
			num /= i
			fmt.Printf("%d ", i)
		}
		if i&0x01 == 0x01 {
			i += 2
		} else {
			i++
		}
	}

	if num > 1 {
		fmt.Print(num)
	}
}
