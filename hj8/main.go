package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	data, _ := reader.ReadString(0)
	lines := strings.Split(data, "\n")
	lines = lines[1:]

	mm := make(map[int]int)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		key, value := keyValue(line)
		mm[key] = mm[key] + value
	}

	keys := make([]int, 0, len(mm))
	for key := range mm {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, key := range keys {
		fmt.Printf("%d %d\n", key, mm[key])
	}
}

func keyValue(str string) (int, int) {
	strs := strings.Split(str, " ")
	key, _ := strconv.ParseInt(strs[0], 10, 64)
	value, _ := strconv.ParseInt(strs[1], 10, 64)
	return int(key), int(value)
}
