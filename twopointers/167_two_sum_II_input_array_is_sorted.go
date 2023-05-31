package twopointers

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left]+numbers[right] > target {
			right -= 1
		} else if numbers[left]+numbers[right] < target {
			left += 1
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return []int{}
}
