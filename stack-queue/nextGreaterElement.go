package stackqueue

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	results := make([]int, len(nums1))
	numToIndex := map[int]int{}

	for _, v := range nums1 {
		numToIndex[v] = -1
	}

	for i, v := range nums2 {
		_, ok := numToIndex[v]
		if ok {
			numToIndex[v] = i
		}

	}

	return results
}
