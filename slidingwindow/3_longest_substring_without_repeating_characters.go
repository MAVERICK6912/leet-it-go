package slidingwindow

func lengthOfLongestSubstring(s string) int {
	charMap := make(map[byte]bool)
	left := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var maxLengthOfLongestUniqueSubstring int
	for index := range s {
		for charMap[s[index]] {
			delete(charMap, s[left])
			left += 1
		}
		charMap[s[index]] = true
		maxLengthOfLongestUniqueSubstring = max(maxLengthOfLongestUniqueSubstring, index-left+1)
	}

	return maxLengthOfLongestUniqueSubstring
}
