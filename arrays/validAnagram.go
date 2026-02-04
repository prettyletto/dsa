package arrays

// first solution
func isAnagram1(s string, t string) bool {
	letters := map[byte]int{}

	for i := range len(s) {
		letters[s[i]]++
	}

	for i := range len(t) {
		if _, ok := letters[t[i]]; !ok || letters[t[i]] <= 0 || len(t) != len(s) {
			return false
		}
		letters[t[i]]--
	}
	return true
}

// better solution
func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var cnt [26]int
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
		cnt[t[i]-'a']--
	}

	for i := range 26 {
		if cnt[i] != 0 {
			return false
		}
	}

	return true
}
