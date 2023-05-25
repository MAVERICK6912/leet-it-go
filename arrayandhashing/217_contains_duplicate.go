package arrayandhashing

func containsDuplicate(nums []int) bool {
	dupMap := make(map[int]int)
	for _, num := range nums {
		dupMap[num] += 1
		if dupMap[num] > 1 {
			return true
		}
	}
	return false
}
