package src

import (
	"fmt"

	"github.com/urfave/cli"
)

func hanoi(c *cli.Context) error {
	hanoiImp(3, "A", "B", "C")
	return nil
}

func hanoiImp(layer int, a, b, c string) {
	if layer <= 1 {
		fmt.Printf("%s -> %s\n", a, c)
	} else {
		hanoiImp(layer-1, a, c, b)
		fmt.Printf("%s -> %s\n", a, c)
		hanoiImp(layer-1, b, a, c)
	}
}
