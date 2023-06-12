package binarysearch

func searchRotated(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if target == nums[mid] {
			return mid
		}
		// check if we're in left sorted portion
		if nums[left] <= nums[mid] {
			if target > nums[mid] || target < nums[left] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			// checking in right sorted portion
			if target < nums[mid] || target > nums[right] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}
