package arrayandhashing

func twoSum(nums []int, target int) []int {
	// diffMap is a map with key as the number and value as index of the element
	diffMap := make(map[int]int)

	for i, num := range nums {
		diff := target - num
		if idx, ok := diffMap[diff]; ok {
			return []int{i, idx}
		}
		diffMap[num] = i
	}
	return []int{}
}
