package stackqueue

type entry struct {
	value int
	min   int
}

type MinStack struct {
	data []entry
}

func Constructor() MinStack {
	return MinStack{data: []entry{}}
}

func (s *MinStack) Push(val int) {
	if len(s.data) == 0 {
		s.data = append(s.data, entry{value: val, min: val})
		return
	}
	currMin := s.data[len(s.data)-1].min
	if val < currMin {
		currMin = val
	}
	s.data = append(s.data, entry{value: val, min: currMin})
}

func (s *MinStack) Pop() {
	if len(s.data) == 0 {
		return
	}
	s.data = s.data[:len(s.data)-1]
}

func (s *MinStack) Top() int {
	return s.data[len(s.data)-1].value
}

func (s *MinStack) GetMin() int {
	return s.data[len(s.data)-1].min
}
