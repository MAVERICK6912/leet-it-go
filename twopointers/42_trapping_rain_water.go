package twopointers

func trap(input []int) int {
	if len(input) == 0 {
		return 0
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	left, right, res := 0, len(input)-1, 0
	maxLeft, maxRight := input[left], input[right]
	for left < right {
		if maxLeft < maxRight {
			left += 1
			maxLeft = max(maxLeft, input[left])
			res += maxLeft - input[left]
		} else {
			right -= 1
			maxRight = max(maxRight, input[right])
			res += maxRight - input[right]
		}
	}
	return res
}
