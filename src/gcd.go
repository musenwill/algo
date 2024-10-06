package src

import (
	"fmt"

	"github.com/urfave/cli"
)

func gcd(c *cli.Context) error {
	gcdCase(1, 1)
	gcdCase(1, 2)
	gcdCase(2, 4)
	gcdCase(3, 7)
	gcdCase(12, 8)
	gcdCase(15, 12)
	return nil
}

func gcdCase(m, n int) {
	cd := gcdImp(m, n)
	fmt.Printf("(%d, %d) -> %d\n", m, n, cd)
}

func gcdImp(m, n int) int {
	if m > n {
		m, n = n, m
	}

	for m > 0 {
		cd := n % m
		if cd > 0 {
			m, n = cd, m
		} else {
			break
		}
	}

	return m
}
