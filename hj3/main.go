package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString(0)
	lines := strings.Split(data, "\n")

	nums := make([]int, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			i, _ := strconv.ParseInt(line, 10, 64)
			nums = append(nums, int(i))
		}
	}

	mm := NewBitMap(500)
	for i := 1; i < len(nums); i++ {
		mm.Put(nums[i])
	}
	mm.Dump()
}

type BitMap struct {
	m []byte
}

func NewBitMap(size int) *BitMap {
	m := make([]byte, size/8+1)
	for i := 0; i < len(m); i++ {
		m[i] = 0
	}
	return &BitMap{m: m}
}

func (b *BitMap) Put(num int) {
	if num < 0 {
		return
	}
	if num >= len(b.m)*8 {
		return
	}

	byteN := num / 8
	bitN := num % 8

	b.m[byteN] |= (1 << bitN)
}

func (b *BitMap) Dump() {
	for i := 0; i < len(b.m)*8; i++ {
		byteN := i / 8
		bitN := i % 8
		mask := byte(1) << bitN

		if b.m[byteN]&mask == mask {
			fmt.Println(i)
		}
	}
}
