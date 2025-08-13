package linkedlist

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.
import "fmt"

// Node represents a single node in the doubly linked list
type Node struct {
	Value interface{}
	next  *Node
	prev  *Node
}

// List represents the doubly linked list structure
type List struct {
	first *Node
	last  *Node
	size  int
}

func NewList(elements ...interface{}) *List {
	newList := &List{
		first: nil,
		last:  nil,
		size:  0,
	}

	for _, element := range elements {
		newNode := &Node{
			Value: element,
			next:  nil,
			prev:  newList.last,
		}

		if newList.last != nil {
			newList.last.next = newNode
		} else {
			newList.first = newNode
		}
		newList.last = newNode
		newList.size++
	}
	return newList
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v interface{}) {
	newNode := &Node{
		Value: v,
		next:  l.first,
		prev:  nil,
	}

	if l.first != nil {
		l.first.prev = newNode
	} else {
		l.last = newNode
	}
	l.first = newNode // point first to new created node.
	l.size++
}

func (l *List) Push(v interface{}) {
	newNode := &Node{
		Value: v,
		next:  nil,
		prev:  l.last,
	}

	if l.last != nil {
		l.last.next = newNode
	} else {
		l.first = newNode
	}
	l.last = newNode
	l.size++
}

func (l *List) Shift() (interface{}, error) {
	if l.size == 0 {
		return nil, fmt.Errorf("Failed")
	}
	result := l.first.Value
	if l.first.next != nil {
		l.first = l.first.next
		l.first.prev = nil
	} else {
		l.first = nil
		l.last = nil
	}
	l.size--
	return result, nil
}

func (l *List) Pop() (interface{}, error) {
	if l.size == 0 {
		return nil, fmt.Errorf("Failed")
	}
	result := l.last.Value
	if l.last.prev != nil {
		l.last = l.last.prev
		l.last.next = nil
	} else {
		l.first = nil
		l.last = nil
	}
	l.size--
	return result, nil
}

func (l *List) Reverse() {
	current := l.first

	for current != nil {
		temp := current.next
		current.next = current.prev
		current.prev = temp
		current = temp
	}

	temp := l.first
	l.first = l.last
	l.last = temp

}

func (l *List) First() *Node {
	return l.first
}

func (l *List) Last() *Node {
	return l.last
}
