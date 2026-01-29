package arrays

func lengthOfLongestSubstring(s string) int {
	seen := map[byte]int{}
	maxLength := 0

	l := 0

	for r := range s {
		if idx, ok := seen[s[r]]; ok && idx >= l {
			l = idx + 1
		}
		seen[s[r]] = r
		maxLength = max(maxLength, r-l+1)
	}

	return maxLength
}
