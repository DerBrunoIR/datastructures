package filter

import (
	iter "DerBrunoIR/datastructures/lib/iterator"
)


type FilterFunc[T any] func(T) bool


func MakeFilter[T any](iter iter.Iterator[T], f FilterFunc[T]) *Filter[T] {
	return &Filter[T]{
		iter: iter,
		f: f,
		hasNext: true,
	}
}

type Filter[T any] struct {
	iter iter.Iterator[T]
	f FilterFunc[T]
	next T
	hasNext bool
	valid bool
}

func (f *Filter[T]) Clone() *Filter[T] {
	if cloner, ok := f.iter.(iter.Cloner[iter.Iterator[T]]); ok {
		return MakeFilter(
			cloner.Clone(),
			f.f,
		)
	}
	panic("iterator doesn't implement Cloner interface")
}

func (f *Filter[T]) findNext() {
	for {
		if !f.iter.HasNext() {
			f.hasNext = false
			break 
		}
		cur := f.iter.Next()
		if f.f(cur) {
			f.next = cur
			break 
		}
	}
}

func (f *Filter[T]) update() {
	if !f.valid {
		f.findNext()
		f.valid = true
	}
}

func (f *Filter[T]) HasNext() bool {
	f.update()
	return f.hasNext
}

func (f *Filter[T]) Next() T {
	if !f.HasNext() {
		panic("no more elments")
	}
	f.valid = false
	return f.next
}

