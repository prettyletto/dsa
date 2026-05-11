package stackqueue

type Deque struct {
	data []int
	head int
	tail int
	size int
}

func NewDeque(capacity int) *Deque {
	if capacity <= 0 {
		capacity = 1
	}

	return &Deque{
		data: make([]int, capacity),
		head: 0,
		tail: 0,
		size: 0,
	}
}

func (d *Deque) PushBack(x int) {
	if d.IsFull() {
		d.grow()
	}

	d.data[d.tail] = x
	d.tail = d.next(d.tail)
	d.size++

}

func (d *Deque) PushFront(x int) {
	if d.IsFull() {
		d.grow()
	}

	d.head = d.prev(d.head)
	d.data[d.head] = x
	d.size++
}

func (d *Deque) PopFront() (int, bool) {
	if d.IsEmpty() {
		return 0, false
	}
	var value int

	value = d.data[d.head]
	d.head = d.next(d.head)
	d.size--

	return value, true
}

func (d *Deque) PopBack() (int, bool) {
	if d.IsEmpty() {
		return 0, false
	}
	var value int

	d.tail = d.prev(d.tail)
	value = d.data[d.tail]
	d.size--
	return value, true
}

func (d *Deque) Front() (int, bool) {
	if d.IsEmpty() {
		return 0, false
	}

	return d.data[d.head], true
}

func (d *Deque) Back() (int, bool) {
	if d.IsEmpty() {
		return 0, false
	}

	idx := d.prev(d.tail)
	return d.data[idx], true
}

func (d *Deque) IsEmpty() bool {
	return d.size == 0
}

func (d *Deque) IsFull() bool {
	return len(d.data) == d.size
}

func (d *Deque) Len() int {
	return d.size
}

func (d *Deque) next(i int) int {
	return (i + 1) % len(d.data)
}

func (d *Deque) prev(i int) int {
	return (i - 1 + len(d.data)) % len(d.data)
}
func (d *Deque) grow() {
	old := d.data
	newData := make([]int, len(old)*2)
	idx := d.head

	for i := 0; i < d.size; i++ {
		newData[i] = old[idx]
		idx = (idx + 1) % len(old)
	}

}
