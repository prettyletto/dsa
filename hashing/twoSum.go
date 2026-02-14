package hashing

func twoSum (nums []int, target int) []int  {
  candidates := make(map[int]int, len(nums))

	for i := range nums {
		if _,ok := candidates[target-nums[i]]; ok {
			return []int {candidates[target-nums[i]], i}
		}
		candidates[nums[i]] = i
	}

	return []int{}
}

