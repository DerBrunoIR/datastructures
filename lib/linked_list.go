package lib

// default, next and prev 
type ListNode interface {
	Next() ListNode 
	Prev() ListNode 
	SetNext(ListNode)
	SetPrev(ListNode)
}

type List[T any] interface {
	Queue[T]
	Stack[T]
	Sizer
	Remove(T)
}

func MakeDoubleLinkedList[T ListNode]() *DoubleLinkedList[T] {
	var head, tail T
	head.SetNext(tail)
	tail.SetPrev(head)
	return &DoubleLinkedList[T]{

	}
}

type DoubleLinkedList[T ListNode] struct {
	head, tail T
	size int
}

func (l *DoubleLinkedList[T]) Push(t T) {
	
}

func (l *DoubleLinkedList[T]) Pop(t T) {}
func (l *DoubleLinkedList[T]) Enqueue(t T) {}
func (l *DoubleLinkedList[T]) Dequeue(t T) {}
func (l *DoubleLinkedList[T]) Size(t T) {}

