package stack_test

import (
	"testing"

	"github.com/supermarine1377/algorithm-go/lib/stack"
)

func Test_Stack(t *testing.T) {
	stack := stack.New()
	if !stack.IsEmpty() {
		t.Error("error: the stack is expected to be empty, but it has some elements;")
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if top := stack.Top(); top != 3 {
		t.Errorf("error: the the element must be 3, but is is %s", top)
	}
	if popped := stack.Pop(); popped != 3 {
		t.Errorf("error: the popped element must be 3, but is is %s", popped)
	}
	if popped := stack.Pop(); popped != 2 {
		t.Errorf("error: the popped element must be 2, but is is %s", popped)
	}
	if popped := stack.Pop(); popped != 1 {
		t.Errorf("error: the popped element must be 1, but is is %s", popped)
	}
	if !stack.IsEmpty() {
		t.Error("error: the stack is expected to be empty, but it has some elements;")
	}
}
