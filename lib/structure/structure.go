package structure


type Sizer interface {
	Size() int 
}

type Capaciter interface {
	Capacity() int
}

type Stack[T any] interface {
	Pop() T
	Push(T)
}

type Queue[T any] interface {
	Enqueue(T)
	Dequeue() T
}
