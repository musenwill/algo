package util

func Assert(condition bool) {
	if !condition {
		panic("")
	}
}
