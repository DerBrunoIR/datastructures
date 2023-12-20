package lib


func MakeBufferedIterator[T any](i Iterator[T], capacity int) *BufferedIterator[T] {
	return &BufferedIterator[T]{
		iter: i,
		buffer: MakeRingBuffer[T](capacity),
	}
}

type BufferedIterator[T any] struct {
	iter Iterator[T]
	buffer Buffer[T]
}

func (i *BufferedIterator[T]) Clone() *BufferedIterator[T] {
	if cloner, ok := i.iter.(Cloner[Iterator[T]]); ok {
		return &BufferedIterator[T]{
			iter: cloner.Clone(),
			buffer: i.buffer.Clone(),
		}
	}
	panic("iterator doesn't implement Cloner interface")
}

func (i *BufferedIterator[T]) fillBuffer() {
	for i.buffer.Capacity() > 0 && i.HasNext() {
		i.buffer.Enqueue(i.Next())
	}
}

func (i *BufferedIterator[T]) HasNext() bool {
	return i.buffer.Size() > 0 || i.HasNext()
}

func (i *BufferedIterator[T]) Next() T {
	if !i.HasNext() {
		panic("no more elements")
	}
	return i.buffer.Dequeue()
}


