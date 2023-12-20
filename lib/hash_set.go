package lib

import "slices"


func MakeHashSet[T Id](capacity int) *HashSet[T] {
	return &HashSet[T]{
		elements: make([][]T, capacity),
		count: 0,
	}
}


type HashSet[T Id] struct {
	elements [][]T
	count int
}

func (h *HashSet[T]) Size() int {
	return h.count
}

func (h *HashSet[T]) lookup(s []T, t T) (int, bool) {
	idx, found := slices.BinarySearchFunc(
		s,
		t,
		func(a, b T) int {
			return a.Id() - b.Id() 
		},
	)
	if !found {
		return -1, false 
	}
	return idx, true
}


func (h *HashSet[T]) Insert(t T) {
	i := t.Id() % len(h.elements)
	if h.elements[i] != nil {
		idx, found := h.lookup(h.elements[i], t)
		if !found {
			h.elements[i] = slices.Insert(h.elements[i], idx, t)
			h.count++
		}
	} else {
		h.elements[i] = []T{t}
	}
}

func (h *HashSet[T]) Remove(t T) {
	i := t.Id() % len(h.elements)
	idx, found := h.lookup(h.elements[i], t)
	if found {
		h.elements[i] = slices.Delete(h.elements[i], idx, idx+1)
		h.count--
	}
}

