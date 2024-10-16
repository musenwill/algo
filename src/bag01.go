package src

import (
	"fmt"

	"github.com/urfave/cli"
)

func bag01(c *cli.Context) error {
	bag01Case([]int{3, 2, 5, 7}, 13)
	bag01Case([]int{3, 2, 5, 7}, 14)
	bag01Case([]int{3, 2, 5, 7}, 15)
	bag01Case([]int{3, 2, 5, 7}, 16)
	return nil
}

func bag01Case(bricks []int, cap int) {
	fmt.Printf("%d = %d-%v\n", bag01Imp(make(map[string]int), bricks, cap), cap, bricks)
}

func bag01Imp(dict map[string]int, bricks []int, cap int) int {
	key := fmt.Sprintf("%d-%d", len(bricks), cap)
	if weight, ok := dict[key]; ok {
		return weight
	}

	if len(bricks) == 0 || cap <= 0 {
		dict[key] = 0
		return 0
	}

	lastBrick := bricks[len(bricks)-1]

	kn1a1 := fmt.Sprintf("%d-%d", len(bricks)-1, cap-lastBrick)
	if _, ok := dict[kn1a1]; !ok {
		bag01Imp(dict, bricks[:len(bricks)-1], cap-lastBrick)
	}
	n1a1 := dict[kn1a1]

	if n1a1+lastBrick <= cap {
		n1a1 += lastBrick
	}

	kna1 := fmt.Sprintf("%d-%d", len(bricks)-1, cap)
	if _, ok := dict[kna1]; !ok {
		bag01Imp(dict, bricks[:len(bricks)-1], cap)
	}
	na1 := dict[kna1]

	dict[key] = max(n1a1, na1)
	return dict[key]
}
