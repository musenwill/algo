package util

type NumT interface {
	int | uint | int64 | uint64 | float32 | float64
}

type Heap[T NumT] struct {
	size int
	buf  []T
	less func(i, j T) bool
}

func NewHeap[T NumT](cap int, less func(i, j T) bool) *Heap[T] {
	hp := &Heap[T]{
		less: less,
		buf:  make([]T, 0, cap+1),
	}
	hp.buf = append(hp.buf, 0)
	return hp
}

func (h *Heap[T]) Push(n T) {
	if h.size == 0 {
		h.buf = append(h.buf, n)
		h.size++
		return
	}
	if h.less(n, h.Peek()) {
		return
	}

	h.size++
	h.buf = append(h.buf, n)
	for parent := h.size >> 1; parent > 0; parent = parent >> 1 {
		if h.less(n, h.buf[parent]) {
			h.buf[h.size], h.buf[parent] = h.buf[parent], h.buf[h.size]
		}
	}
}

func (h *Heap[T]) BatchPush(n ...T) {
	for _, i := range n {
		h.Push(i)
	}
}

func (h *Heap[T]) Peek() T {
	if h.size <= 0 {
		return 0
	}
	return h.buf[1]
}

func (h *Heap[T]) Size() int {
	return h.size
}

func (h *Heap[T]) Gets() []T {
	return h.buf[1 : h.size+1]
}
