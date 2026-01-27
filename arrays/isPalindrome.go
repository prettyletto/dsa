package arrays

import (
	"regexp"
	"strings"
)

func isAlnum(b byte) bool {
	return ( b >= 'a' && b <= 'z') || (b>= '0' && b <= '9')
}


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
		l ++ 
		r --
	}
	return true
}

