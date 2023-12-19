package buffers 


func MakeRingBuffer[T any](size int) *RingBuffer[T] {
	return &RingBuffer[T]{
		buffer: make([]T, size),
	}
}

type RingBuffer[T any] struct {
	buffer []T
	start, end int
}


