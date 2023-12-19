package lib

import "slices"

func MakeRingBuffer[T any](capacity int) *RingBuffer[T] {
	if capacity < 1 {
		panic("ring buffer capacity musst be greater zero")
	}
	return &RingBuffer[T]{
		buffer: make([]T, capacity+1),
	}
}

/*
This type implements a circulare buffer
*/
type RingBuffer[T any] struct {
	buffer     []T
	start, end int
}

func (r *RingBuffer[T]) Size() int {
	return (r.end - r.start) % len(r.buffer)
}

func (r *RingBuffer[T]) Capacity() int {
	return (r.start - r.end - 1) % len(r.buffer)
}

func (r *RingBuffer[T]) Enqueue(t T) {
	if r.end+1 == r.start {
		panic("no more space")
	}
	r.buffer[r.end] = t
	r.end++
	r.end %= len(r.buffer)
}

func (r *RingBuffer[T]) Dequeue() (res T) {
	if r.start == r.end {
		panic("no more elements")
	}
	res = r.buffer[r.start]
	r.start++
	r.start %= len(r.buffer)
	return
}

func (r *RingBuffer[T]) Clone() *RingBuffer[T] {
	return &RingBuffer[T]{
		buffer: slices.Clone(r.buffer),
		start: r.start,
		end: r.end,
	}
}
