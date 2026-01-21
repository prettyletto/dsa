package arrays

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums))

	sum := 0
	for i := range nums {
		sum += nums[i]
		sums[i] = sum
	}

	return NumArray{sums: sums}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return this.sums[right]
	}

	return this.sums[right] - this.sums[left-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
