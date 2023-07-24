package Q54

type MyStack struct {
	stack []int
}

func Constructor() MyStack {
	return MyStack{stack: []int{}}
}

func (this *MyStack) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MyStack) Pop() int {
	arrLen := len(this.stack)
	top := this.stack[arrLen-1]
	// 切掉最後一個元素
	this.stack = this.stack[:arrLen-1]
	return top
}

func (this *MyStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.stack) == 0
}
