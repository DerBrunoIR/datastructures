package trees 



type Element interface {
	Key() int 
}

type Heap interface {
	Insert(Element)
	Remove(Element)
	Extract() Element 
	Get() Element
}



type BinaryTree interface {
	Insert(Element)
	Remove(Element)
	InOrder() Iterator[Element]
	PreOrder() Iterator[Element]
	OutOrder() Iterator[Element]
}

type ArrayTree[T any] struct {
	elements []T
	
}

type BinaryMinHeap struct {
}
