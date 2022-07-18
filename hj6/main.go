package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	_, _ = fmt.Scan(&num)
	prims := NewPrims(num)
	prim := prims.Next()

	for num > 1 {
		if num%prim == 0 {
			num /= prim
			fmt.Printf("%d ", prim)
		} else {
			prim = prims.Next()
			fmt.Println(prim)
		}
	}
}

type Prims struct {
	prims []int
	index int
	max   int
}

func NewPrims(max int) *Prims {
	prims := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41}
	return &Prims{
		prims: prims,
		index: -1,
		max:   max,
	}
}

func (p *Prims) Next() int {
	p.index++
	if p.index < len(p.prims) {
		return p.prims[p.index]
	}
	return p.nexPrim()
}

func (p *Prims) nexPrim() int {
	for num := p.prims[len(p.prims)-1] + 2; num <= p.max; num += 2 {
		if p.isPrim(num) {
			p.prims = append(p.prims, num)
			return num
		} else {
			continue
		}
	}
	return 0
}

func (p *Prims) isPrim(num int) bool {
	for _, prim := range p.prims {
		if prim > int(math.Sqrt(float64(num))) {
			break
		}
		if num%prim == 0 {
			return false
		}
	}
	return true
}
