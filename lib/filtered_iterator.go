package lib


type FilterIteratorFunc[T any] func(T) bool


func MakeFilterIterator[T any](iter Iterator[T], f FilterIteratorFunc[T]) *FilterIterator[T] {
	return &FilterIterator[T]{
		iter: iter,
		f: f,
		hasNext: true,
	}
}

type FilterIterator[T any] struct {
	iter Iterator[T]
	f FilterIteratorFunc[T]
	next T
	hasNext bool
	valid bool
}

func (f *FilterIterator[T]) Clone() *FilterIterator[T] {
	if cloner, ok := f.iter.(Cloner[Iterator[T]]); ok {
		return MakeFilterIterator(
			cloner.Clone(),
			f.f,
		)
	}
	panic("iterator doesn't implement Cloner interface")
}

func (f *FilterIterator[T]) findNext() {
	for {
		if !f.HasNext() {
			f.hasNext = false
			break 
		}
		cur := f.Next()
		if f.f(cur) {
			f.next = cur
			break 
		}
	}
}

func (f *FilterIterator[T]) update() {
	if !f.valid {
		f.findNext()
		f.valid = true
	}
}

func (f *FilterIterator[T]) HasNext() bool {
	f.update()
	return f.hasNext
}

func (f *FilterIterator[T]) Next() T {
	if !f.HasNext() {
		panic("no more elments")
	}
	f.valid = false
	return f.next
}

