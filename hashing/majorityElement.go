package hashing

func majorityElement(nums []int) int {
	counter := make(map[int]int, len(nums)) 
	
	for _, v := range nums {
		counter[v]++
		if counter[v] > len(nums) /2 {
			return v
		}
	}
	return 0
}

func majorityElement2(nums []int) int {
	candidate := 0
	count := 0

	for _,v := range nums {
		if count == 0 {
			candidate = v
		}

		if v == candidate {
			count ++ 
		} else {
			count --
		}
	}

	return candidate
}
