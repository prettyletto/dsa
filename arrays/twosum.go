package arrays


func twoSum(nums []int, target int) []int {
	for i :=  range  nums{
		cdt := nums[i]
		for j:= i + 1; j < len(nums); j++ {
			if cdt + nums[j] == target {
				return []int{i,j}
			}
		}
	}
	return []int{}
}


