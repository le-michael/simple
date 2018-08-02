package simple

import "errors"

type (
	stackNode struct {
		value interface{}
		next  *stackNode
	}

	//Stack struct
	Stack struct {
		top *stackNode
	}
)

//Empty returns if the stack is empty as a boolean
func (stack *Stack) Empty() bool {
	return stack.top == nil
}

//Top returns the value of the node at the top of the stack
func (stack *Stack) Top() interface{} {
	if stack.Empty() {
		return nil
	}
	return stack.top.value
}

//Push a value to the top of the stack
func (stack *Stack) Push(data interface{}) {
	stack.top = &stackNode{data, stack.top}
}

//Pop a value from the stack
func (stack *Stack) Pop() error {
	if stack.Empty() {
		return errors.New("stack: Cannot Pop from empty stack")
	}
	stack.top = stack.top.next
	return nil
}

//NewStack creates a new Stack object
func NewStack() *Stack {
	return &Stack{nil}
}
