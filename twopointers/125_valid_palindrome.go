package twopointers

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isAlphaNumeric(s[left]) {
			left += 1
		}
		for right > left && !isAlphaNumeric(s[right]) {
			right -= 1
		}
		if !isLowerEqual(s[left], s[right]) {
			return false
		}
		left, right = left+1, right-1
	}
	return true
}

func isAlphaNumeric(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}

func isLowerEqual(l, r byte) bool {
	if l >= 'A' && l <= 'Z' {
		if r >= 'A' && r <= 'Z' {
			return l+32 == r+32
		}
		return l+32 == r
	} else {
		if r >= 'A' && r <= 'Z' {
			return l == r+32
		}
		return l == r
	}
}
