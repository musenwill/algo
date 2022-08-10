package numbers

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"time"
	"unsafe"
)

type Generator struct {
	total int
	count int
	cap   int
	file  string
	buf   []int64
}

func NewGenerator(total int, file string) *Generator {
	rand.Seed(time.Now().Unix())
	cap := 1024 * 1024
	if total < cap {
		cap = total
	}
	return &Generator{
		file:  file,
		total: total,
		cap:   cap,
		buf:   make([]int64, 0, cap),
	}
}

func (g *Generator) GenNumbers() error {
	fd, err := os.OpenFile(g.file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("failed open file, %w", err)
	}
	defer fd.Close()

	for g.count < g.total {
		bytes := g.GenSlice()
		if _, err := fd.Write(bytes); err != nil {
			return err
		}
		g.buf = g.buf[:0]
	}

	return nil
}

func (g *Generator) GenSlice() []byte {
	for i := 0; i < g.cap; i++ {
		g.count++
		g.buf = append(g.buf, rand.Int63())
	}
	var bytes []byte
	originHeader := (*reflect.SliceHeader)(unsafe.Pointer(&g.buf))
	header := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	header.Data = originHeader.Data
	header.Len = originHeader.Len * 8
	header.Cap = originHeader.Cap * 8
	return bytes
}

func LoadSlice(data []byte) []int64 {
	var buf []int64
	originHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	header := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	header.Data = originHeader.Data
	header.Len = originHeader.Len / 8
	header.Cap = originHeader.Cap / 8
	return buf
}
