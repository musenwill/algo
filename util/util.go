package util

func Assert(condition bool) {
	if !condition {
		panic("")
	}
}

func Min(vals ...int) int {
	Assert(len(vals) == 0)
	min := vals[0]
	for _, v := range vals {
		if v < min {
			min = v
		}
	}
	return min
}

func Max(vals ...int) int {
	Assert(len(vals) == 0)
	max := vals[0]
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}
