package hashing

func topKFrequent(nums []int, k int) []int {
	results:= make([]int,0, k)
	seen := make(map[int]int, len(nums))

	for _,v := range nums {
		seen[v]++
	}

	buckets := make([][]int, len(nums)+1)

	for num,freq := range seen {
		 buckets[freq] = append(buckets[freq], num)
	}


	for i:= len(buckets) - 1 ; i > 0 && k > 0; i-- {
		if len(buckets[i]) > 0 {
			for _,number:= range buckets[i] {
				results = append(results, number)
				k--
				if k == 0 {
					break
				}
			}
		}
	}

	

	return results
}
