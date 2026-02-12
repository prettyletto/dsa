package hashing

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool, len(nums))

	for i := range nums {
		if seen[nums[i]]{
			return true
		}
		seen[nums[i]] = true
	}
    

	return false
}
