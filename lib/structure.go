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

type Set[T Id] interface {
	Contains(T) bool
	Insert(T)
	Remove(T)
}

type Tree[T Id] interface {
	
}

type TreeNode[T any] interface {
	Value() T
}

