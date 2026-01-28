package arrays

import (
	"regexp"
	"strings"
)

func keepLettersAndDigitsRegex(s string) string {
	reg := regexp.MustCompile(`[^\p{L}\p{N}]+`)
	cleaned := reg.ReplaceAllString(s, "")
	return cleaned
}

func isPalindrome(s string) bool {
	cleaned := strings.ToLower(keepLettersAndDigitsRegex(s))
	l := 0
	r := len(cleaned) - 1

	for l <= r {
		if cleaned[l] != cleaned[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlnum(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9')
}

func isPalindromePointer(s string) bool {
	s = strings.ToLower(s)

	l, r := 0, len(s)-1

	for l < r {
		for l < r && !isAlnum(s[l]) {
			l++
		}
		for l < r && !isAlnum(s[r]) {
			r--
		}

		if s[l] != s[r] {
			return false
		}
		l--
		r++
	}
	return true
}
