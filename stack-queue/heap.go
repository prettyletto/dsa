package stackqueue

type MyQueue struct {
	data []int
}

func Constructor() MyQueue {
	return MyQueue{data: []int{}}
}

func (this *MyQueue) Push(x int) {
	this.data = append(this.data, x)
}

func (this *MyQueue) Pop() int {
	var value int
	if this.Empty() {
		return value
	}
	value = this.data[0]
	this.data = this.data[1:]
	return value
}

func (this *MyQueue) Peek() int {
	var value int
	if this.Empty() {
		return value
	}
	value = this.data[0]
	return value
}

func (this *MyQueue) Empty() bool {
	return len(this.data) <= 0
}
