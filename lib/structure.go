package lib


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

type Id interface {
	Id() int
}

type Container[T any] interface {
	Contains(T) bool 
}

type Set[T Id] interface {
	Container[T]
	Insert(T)
	Remove(T)
}

