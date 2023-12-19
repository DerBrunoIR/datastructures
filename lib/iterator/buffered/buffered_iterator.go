package buffered

import (
	iter "DerBrunoIR/datastructures/lib/iterator"
	ring "DerBrunoIR/datastructures/lib/structure/ring"
)



func Make[T any](i iter.Iterator[T], capacity int) *Buffered[T] {
	return &Buffered[T]{
		iter: i,
		buffer: *ring.Make[T](capacity),
	}
}

type Buffered[T any] struct {
	iter iter.Iterator[T]
	buffer ring.Buffer[T]
}

func (i *Buffered[T]) Clone() *Buffered[T] {
	if cloner, ok := i.iter.(iter.Cloner[iter.Iterator[T]]); ok {
		return &Buffered[T]{
			iter: cloner.Clone(),
			buffer: *i.buffer.Clone(),
		}
	}
	panic("iterator doesn't implement Cloner interface")
}

func (i *Buffered[T]) fillBuffer() {
	for i.buffer.Capacity() > 0 && i.iter.HasNext() {
		i.buffer.Enqueue(i.iter.Next())
	}
}

func (i *Buffered[T]) HasNext() bool {
	return i.buffer.Size() > 0 || i.iter.HasNext()
}

func (i *Buffered[T]) Next() T {
	if !i.HasNext() {
		panic("no more elements")
	}
	return i.buffer.Dequeue()
}


