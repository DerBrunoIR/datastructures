package lib


func MakeBufferedIterator[T any](i Iterator[T], capacity int) *BufferedIterator[T] {
	return &BufferedIterator[T]{
		iter: i,
		buffer: *MakeRingBuffer[T](capacity),
	}
}

type BufferedIterator[T any] struct {
	iter Iterator[T]
	buffer RingBuffer[T]
}

func (i *BufferedIterator[T]) fillBuffer() {
	for i.buffer.Capacity() > 0 && i.iter.HasNext() {
		i.buffer.Enqueue(i.iter.Next())
	}
}

func (i *BufferedIterator[T]) HasNext() bool {
	return i.buffer.Size() > 0 || i.iter.HasNext()
}

func (i *BufferedIterator[T]) Next() T {
	if !i.HasNext() {
		panic("no more elements")
	}
	return i.buffer.Dequeue()
}


