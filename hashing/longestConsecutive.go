package hashing

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	seen := make(map[int]bool, len(nums))
	for _, n := range nums {
		seen[n] = true
	}

	best := 0
	for n := range seen {
		if seen[n-1] {
			continue
		}

		length := 1
		for seen[n+length] {
			length++
		}
		if length > best {
			best = length
		}
	}

	return best
}
