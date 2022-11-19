package stack

import "container/list"

type Stack struct {
	list *list.List
}

func New() *Stack {
	list := list.New()
	return &Stack{list: list}
}

// Push() performs pushing to the stack, and returns nil if the operation is successful and nil if the operation fails
func (s *Stack) Push(element interface{}) {
	s.list.PushBack(element)
}

// Top() returns a top element of the stack
func (s *Stack) Top() interface{} {
	element := s.list.Back()
	return element.Value
}

// Pop() removes a top element from the stack
func (s *Stack) Pop() interface{} {
	back := s.list.Back()
	if back == nil {
		return nil
	}
	return s.list.Remove(back)
}

// IsEmpty() returns bool which shows if the stack has any elements
func (s *Stack) IsEmpty() bool {
	return s.list.Len() == 0
}
