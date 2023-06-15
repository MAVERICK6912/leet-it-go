package slidingwindow

func characterReplacement(s string, k int) int {
	// We can also use a  count := make([]int, 91) and it will do exact same Job ("Z"'s ASCII value is 90)
	count := make(map[byte]int)
	res, maxf, left := 0, 0, 0

	for right := 0; right < len(s); right++ {
		count[s[right]] += 1 // s[right] will return a byte. Incase we used slice instead of hashmap this byte is the same as that element's position in the count slice

		maxf = max(maxf, count[s[right]])

		// Move left pointer
		if (right-left+1)-maxf > k {
			count[s[left]] -= 1 // When we decrement, we dont have to update max variable as reducing max variable wont have any effect on end result
			left += 1
		}

		res = max(res, right-left+1)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
