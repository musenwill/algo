package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _, _ := reader.ReadLine()

	runes := []rune(string(input))

	exprStack := &stack{d: make([]rune, 0)}
	opStack := &stack{d: make([]rune, 0)}

	for _, r := range runes {
		if isDigital(r) {
			exprStack.push(r)
			continue
		}
		switch r {
		case '(':
			opStack.push(r)
		case ')':
			for o := opStack.pop(); o != '('; o = opStack.pop() {
				exprStack.push(o)
			}
		default:
			if opStack.len() == 0 {
				opStack.push(r)
			} else {
				if priority[r] >= priority[opStack.peak()] {
					opStack.push(r)
				} else {
					for opStack.len() > 0 && priority[opStack.peak()] > priority[r] {
						exprStack.push(opStack.pop())
					}
					opStack.push(r)
				}
			}
		}
	}
	for opStack.len() > 0 {
		exprStack.push(opStack.pop())
	}

	resultStack := &resultStack{d: make([]float64, 0)}
	for _, r := range exprStack.d {
		if isDigital(r) {
			i, _ := strconv.ParseInt(string(r), 10, 64)
			resultStack.push(float64(i))
		} else {
			b := resultStack.pop()
			a := resultStack.pop()
			resultStack.push(calc(a, b, r))
		}
	}
	
	fmt.Println(resultStack.pop())
}

var priority = map[rune]int {
	'+': 1,
	'-': 1,
	'*': 2,
	'/': 2,
}

type stack struct {
	d []rune
}

func (s *stack) len() int {
	return len(s.d)
}

func (s *stack) push(r rune) {
	s.d = append(s.d, r)
}

func (s *stack) peak() rune {
	if s.len() > 0 {
		return s.d[s.len()-1]
	}
	return 0
}

func (s *stack) pop() rune {
	if s.len() <= 0 {
		return 0
	}
	v := s.d[s.len()-1]
	s.d = s.d[:s.len()-1]
	return v
}

func isDigital(r rune) bool {
	return r >= '0' && r <= '9'
}


type resultStack struct {
	d []float64
}

func (s *resultStack) len() int {
	return len(s.d)
}

func (s *resultStack) push(r float64) {
	s.d = append(s.d, r)
}

func (s *resultStack) peak() float64 {
	if s.len() > 0 {
		return s.d[s.len()-1]
	}
	return 0
}

func (s *resultStack) pop() float64 {
	if s.len() <= 0 {
		return 0
	}
	v := s.d[s.len()-1]
	s.d = s.d[:s.len()-1]
	return v
}

func calc(a, b float64, op rune) float64 {
	switch op {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		return a / b
	}
	return 0
}