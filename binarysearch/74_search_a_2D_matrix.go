package binarysearch

func searchMatrix(matrix [][]int, target int) bool {
	numberOfRows, columns := len(matrix), len(matrix[0])
	topRow, bottomRow := 0, numberOfRows-1
	for topRow <= bottomRow {
		midRow := topRow + (bottomRow-topRow)/2
		if matrix[midRow][columns-1] < target {
			topRow = midRow + 1
		} else if matrix[midRow][0] > target {
			bottomRow = midRow - 1
		} else {
			break
		}
	}
	if !(topRow <= bottomRow) {
		return false
	}
	rowToSearch := topRow + (bottomRow-topRow)/2
	left, right := 0, columns-1

	for left <= right {
		mid := left + (right-left)/2

		if matrix[rowToSearch][mid] > target {
			right = mid - 1
		} else if matrix[rowToSearch][mid] < target {
			left = mid + 1
		} else {
			return true
		}
	}

	return false
}
