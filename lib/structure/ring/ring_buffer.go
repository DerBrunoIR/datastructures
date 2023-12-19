package ring

import "slices"

func Make[T any](capacity int) *Buffer[T] {
	if capacity < 1 {
		panic("ring buffer capacity musst be greater zero")
	}
	return &Buffer[T]{
		buffer: make([]T, capacity+1),
	}
}

/*
This type implements a circulare buffer
*/
type Buffer[T any] struct {
	buffer     []T
	start, end int
}

func (r *Buffer[T]) Size() int {
	return (r.end - r.start) % len(r.buffer)
}

func (r *Buffer[T]) Capacity() int {
	return (r.start - r.end - 1) % len(r.buffer)
}

func (r *Buffer[T]) Enqueue(t T) {
	if r.end+1 == r.start {
		panic("no more space")
	}
	r.buffer[r.end] = t
	r.end++
	r.end %= len(r.buffer)
}

func (r *Buffer[T]) Dequeue() (res T) {
	if r.start == r.end {
		panic("no more elements")
	}
	res = r.buffer[r.start]
	r.start++
	r.start %= len(r.buffer)
	return
}

func (r *Buffer[T]) Clone() *Buffer[T] {
	return &Buffer[T]{
		buffer: slices.Clone(r.buffer),
		start: r.start,
		end: r.end,
	}
}
