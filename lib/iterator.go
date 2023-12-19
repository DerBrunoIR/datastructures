package iterator


type Cloner[T any] interface {
	Clone() T
}

type Iterator[T any] interface {
	Next() T
	HasNext() bool
}

// Allocates a new slice and appends all elements from the iterator.
func Make[T any](i Iterator[T]) []T {
	s := make([]T, 0)
	return Append(i, s)
}

// Append all elements to the given slice.
func Append[T any](i Iterator[T], s []T) []T {
	for i.HasNext() {
		s = append(s, i.Next())
	}
	return s
}

