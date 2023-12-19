package iterator


type Iterator[T any] interface {
	Next() T
	HasNext() bool
}

func Append[T any](i Iterator[T], s []T) []T {
	for i.HasNext() {
		s = append(s, i.Next())
	}
	return s
}
