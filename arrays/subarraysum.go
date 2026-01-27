package arrays

func subarraySumBruteForce(nums []int, k int) int {
	total := 0
	for i := range nums {
		currSum := 0
		for j := i; j < len(nums); j++ {
			currSum += nums[j]
			if currSum == k {
				total++
			}
		}
	}
	return total
}

func subarraySum(nums []int, k int) int {
	total := 0
	currSum := 0
	prefix := map[int]int{0: 1}
	for i := range nums {
		currSum += nums[i]
		v, ok := prefix[currSum-k]
		if ok {
			total += v
		}
		prefix[currSum]++
	}

	return total
}
