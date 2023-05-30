package arrayandhashing

func longestConsecutive(nums []int) int {
	numberSet := make(map[int]bool)
	for _, num := range nums {
		numberSet[num] = true
	}
	// initialize longest length of sub-sequence
	longest := 0
	for _, num := range nums {
		if !numberSet[num-1] {
			// initialize length of current consecutive sub-sequence
			length := 1
			// loop till num+length is present in numberSet
			for numberSet[num+length] {
				length += 1
			}
			longest = max(length, longest)
		}
	}
	return longest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
