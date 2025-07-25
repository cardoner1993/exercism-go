package linkedlist

import (
	"errors"
)

// Define the List and Element types here.
type Element struct {
	Data interface{}
	Next *Element
}

type List struct {
	Head *Element
	size int
}

func New(elements []int) *List {
	var lst = List{
		Head: nil,
		size: 0,
	}

	for _, elem := range elements {
		// Inline append logic instead of calling method
		newElement := &Element{Data: elem, Next: nil}

		if lst.Head == nil {
			lst.Head = newElement
		} else {
			current := lst.Head
			for current.Next != nil {
				current = current.Next
			}
			current.Next = newElement
		}
		lst.size++
	}

	return &lst
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	// Inline append logic instead of calling method
	newElement := &Element{Data: element, Next: nil}

	if l.Head == nil {
		l.Head = newElement
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newElement
	}
	l.size++
}

func (l *List) Pop() (int, error) {
	if l.Head == nil {
		return 0, errors.New("list is empty")
	}

	// Special case: only one element
	if l.Head.Next == nil {
		result := l.Head.Data.(int)
		l.Head = nil
		l.size--
		return result, nil
	}

	current := l.Head
	for current.Next.Next != nil {
		current = current.Next
	}
	result := current.Next.Data.(int)
	current.Next = nil
	l.size--
	return result, nil
}

func (l *List) Array() []int {
	result := []int{}
	if l.Head == nil {
		return result
	} else {
		current := l.Head
		for current.Next != nil {
			result = append(result, current.Data.(int))
			current = current.Next
		}
		result = append(result, current.Data.(int))
		return result
	}
}

func (l *List) Reverse() *List {
	reversed := &List{Head: nil, size: 0}
	if l.Head == nil {
		return reversed
	}
	current := l.Head
	for current != nil {
		// Insert each element at the beginning of the new list
		newElement := &Element{Data: current.Data, Next: reversed.Head}
		reversed.Head = newElement
		reversed.size++
		current = current.Next
	}
	return reversed
}
