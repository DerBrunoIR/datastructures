package lib


func MakeSliceIterator[S ~[]E, E any](s S, start, count int) *SliceIterator[S, E] {
	return &SliceIterator[S, E]{
		s: s,
		cur: start,
		left: count,
	}
}

type SliceIterator[S ~[]E, E any] struct {
	s S
	cur, left int
}

func (i *SliceIterator[S, E]) Next() E {
	i.cur++
	return i.s[i.cur-1]
}

func (i *SliceIterator[S, E]) HasNext() bool {
	return i.left > 0 
}
