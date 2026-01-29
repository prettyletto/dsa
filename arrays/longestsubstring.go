package arrays 


func lengthOfLongestSubstring( s string)int {
	seen := map[byte]bool{}
	maxLength := 0 

	leng :=  0
	for i := range s {
		if seen[s[i]] {
			maxLength = max(maxLength, leng)
			leng = 0
		}
	}




	return maxLength
}
