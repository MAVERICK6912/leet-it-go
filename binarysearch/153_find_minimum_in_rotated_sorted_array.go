package binarysearch

func findMin(nums []int) int {
	left, right, res := 0, len(nums)-1, nums[0]
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for left <= right {
		if nums[left] < nums[right] {
			res = min(res, nums[left])
			break
		}
		mid := left + (right-left)/2
		res = min(res, nums[mid])
		if nums[mid] >= nums[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}
