package arrayandhashing

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res = fillWithValue(res, 1)
	// find all prefixes
	prefix := 1
	for index, num := range nums {
		res[index] = prefix
		prefix *= num
	}
	// find all postfix and multiply to prefix to get final slice
	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= postfix
		postfix *= nums[i]
	}
	return res
}

func fillWithValue[T any](slice []T, value T) []T {
	for index := range slice {
		slice[index] = value
	}
	return slice
}
