package twopointers

func maxArea(height []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	res := 0
	left, right := 0, len(height)-1
	for left < right {
		area := (right - left) * min(height[left], height[right])
		res = max(res, area)

		if height[left] < height[right] {
			left += 1
		} else {
			right -= 1
		}
	}
	return res
}
