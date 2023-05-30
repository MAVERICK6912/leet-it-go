package arrayandhashing

func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num] += 1
	}
	freqOfElements := make([][]int, len(nums)+1)

	for num, count := range countMap {
		freqOfElements[count] = append(freqOfElements[count], num)
	}
	var res []int
	for i := len(nums); i > 0; i -= 1 {
		for _, v := range freqOfElements[i] {
			res = append(res, v)
			if len(res) == k {
				return res
			}
		}
	}
	return []int{}
}
