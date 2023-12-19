package iterator


func MakeBufferedIterator[T any](i Iterator[T], capacity int) *BufferedIterator[T] {
	return &BufferedIterator[T]{
		iter: i,
		buffer: ,
		start: 0,
		end: 0,
	}
}

type BufferedIterator[T any] struct {
	iter Iterator[T]
	buffer [buffer_size]T
	start, end int
}

func (i *BufferedIterator[T]) fillBuffer() {
	for j := i.start; i.iter.HasNext() && j + 1 != i.end; j = (j+1) % buffer_size {
		i.buffer[j] = i.iter.Next()
	}
}

func (i *BufferedIterator[T]) HasNext() bool {
	i.fillBuffer()
	return i.start == i.end
}

func (i *BufferedIterator[T]) Next() T {
	if !i.HasNext() {
		panic("no more elements")
	}
	res := i.buffer[i.start]
	i.start++
	return res
}


