// stack.go
package stack

import "errors"

type Stack []interface{}

func (stack Stack) Len() int {
	return len(stack)
}

func (stack Stack) Cap() int {
	return cap(stack)
}

func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}

func (stack *Stack) Push(e interface{}) {
	*stack = append(*stack, e)
}

func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("stack is empty")
	}
	return stack[len(stack)-1], nil
}

func (stack *Stack) Pop() (interface{}, error) {
	stk := *stack	// dereference to a local variable stk
	if len(stk) == 0 {
		return nil, errors.New("stack is empty")
	}
	top := stk[len(stk)-1]
	*stack = stk[:len(stk)-1] // shrink the stack
	return top, nil
}
