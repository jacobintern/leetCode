package Q56

type MyQueue struct {
	stack []int
}

func Constructor() MyQueue {
	return MyQueue{stack: []int{}}
}

func (this *MyQueue) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MyQueue) Pop() int {
	head := this.stack[0]
	this.stack = this.stack[1:]
	return head
}

func (this *MyQueue) Peek() int {
	return this.stack[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.stack) == 0
}
