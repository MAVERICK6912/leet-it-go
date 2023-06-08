package stack

type MinStack struct {
	minStack stack[int]
	stack    stack[int]
}

func Constructor() MinStack {
	return MinStack{minStack: make([]int, 0), stack: make(stack[int], 0)}
}

func (this *MinStack) Push(val int) {
	if len(this.minStack) == 0 || val < this.minStack.Peek() {
		this.minStack.Push(val)
	} else {
		this.minStack.Push(this.minStack.Peek())
	}
	this.stack.Push(val)
}

func (this *MinStack) Pop() {
	this.stack.Pop()
	this.minStack.Pop()
}

func (this *MinStack) Top() int {
	return this.stack.Peek()
}

func (this *MinStack) GetMin() int {
	return this.minStack.Peek()
}
