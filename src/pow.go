package src

import (
	"fmt"

	"github.com/urfave/cli"
)

func pow(c *cli.Context) error {
	powCase(1, 1)
	powCase(1, 2)
	powCase(1, 5)
	powCase(2, 0)
	powCase(2, 1)
	powCase(2, 3)
	powCase(2, 5)
	powCase(2, 8)
	powCase(2, 13)
	powCase(2, 16)
	powCase(2, 21)
	return nil
}

func powCase(n, e int) {
	fmt.Printf("%d^%d = %d\n", n, e, powImp(n, e))
}

func powImp(n, e int) int {
	if e == 0 {
		return 1
	}

	r := n
	f := 1
	for e > 1 {
		if 0x01&e == 0x01 {
			f *= r
		}
		r *= r
		e >>= 1
	}

	return r * f
}
