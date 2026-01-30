package arrays

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := range k {
		sum += nums[i]
	}

	maxSum := sum
	l := 0

	for r := k; r < len(nums); r++ {
		sum += nums[r]
		sum -= nums[l]
		l++

		maxSum = max(maxSum, sum)
	}

	return float64(maxSum) / float64(k)
}
