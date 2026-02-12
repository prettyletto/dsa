package arrays


func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		for len(prefix) > len(strs[i]) {
			prefix = prefix[:len(prefix)-1]
		}

		for len(prefix) > 0 && strs[i][:len(prefix)] != prefix {
			prefix = prefix[:len(prefix)-1]
		}

		if prefix == "" {
			return ""
		}
	}

	return prefix
}
