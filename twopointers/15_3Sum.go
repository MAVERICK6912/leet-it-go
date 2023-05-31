package twopointers

import "sort"

func threeSum(nums []int) [][]int {
	// sorting the array to make check for duplicates easier
	// here we'll take O(nlogn) time
	sort.Ints(nums)
	res := make([][]int, 0)
	for index, num := range nums {
		if index > 0 && num == nums[index-1] {
			continue
		}
		left, right := index+1, len(nums)-1
		// this loop will take O(n^2) time
		for left < right {
			threeSum := nums[left] + nums[right] + num
			if threeSum > 0 {
				right -= 1
			} else if threeSum < 0 {
				left += 1
			} else {
				res = append(res, []int{nums[left], nums[right], num})
				left += 1
				for nums[left] == nums[left-1] && left < right {
					left += 1
				}
			}
		}
	}
	return res
}
