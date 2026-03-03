package stackqueue

func dailyTemperatures(temperatures []int) []int {
	results := make([]int, len(temperatures))
	stack := []int{}

	for i, t := range temperatures {
		for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {

			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			results[prevIndex] = i - prevIndex
		}

		stack = append(stack, i)
	}

	return results
}
