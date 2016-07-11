package stack

import (
	"container/list"
	"errors"
)

// Stack is an interface for the Abstract Data Type "Stack"
type Stack interface {
	Push(int)
	Pop() (int, error)
	Peek() (int, error)
}

// SliceStack is a struct with a slice and should
// implement the Stack interface
type SliceStack struct {
	s []int
}

// NewSliceStack should return an empty SliceStack
func NewSliceStack() SliceStack {
	return SliceStack{s: make([]int, 0, 1)}
}

// Push should add an item to the top of the stack
func (ss *SliceStack) Push(item int) {
	ss.s = append(ss.s, item)
}

// Pop should remove an item from the top of the stack
// It returns a tuple with the removed item and an error indicator
// Error will be non-nil if stack is empty
func (ss *SliceStack) Pop() (int, error) {
	if len(ss.s) == 0 {
		return 0, errors.New("empty")
	}
	top := ss.s[len(ss.s)-1]
	ss.s = ss.s[:len(ss.s)-1]
	return top, nil
}

// Peek should inspect item from the top of the stack.
// Like Pop(), but no change made to stack
func (ss *SliceStack) Peek() (int, error) {
	if len(ss.s) == 0 {
		return 0, errors.New("empty")
	}
	return ss.s[len(ss.s)-1], nil
}

// ListStack is an alternate implementation of the Abstract Data Type "Stack"
// using the standard library doubly linked list
type ListStack struct {
	pList *list.List
}

// NewListStack returns a new, empty ListStack
func NewListStack() ListStack {
	return ListStack{pList: list.New()}

}

// Push inserts a value on the stack by putting it on the front of the list
func (ls *ListStack) Push(value int) {
	ls.pList.PushFront(value)
}

// Pop removes the value from the top the stack and returns it
// Error will be non-nil if stack is empty
func (ls *ListStack) Pop() (int, error) {
	frontElement := ls.pList.Front()
	if frontElement == nil {
		return 0, errors.New("empty")
	}
	result := frontElement.Value.(int)
	ls.pList.Remove(frontElement)
	return result, nil
}

// Peek is like Pop, but does not alter the stack
func (ls *ListStack) Peek() (int, error) {
	frontElement := ls.pList.Front()
	if frontElement == nil {
		return 0, errors.New("empty")
	}
	return frontElement.Value.(int), nil
}
