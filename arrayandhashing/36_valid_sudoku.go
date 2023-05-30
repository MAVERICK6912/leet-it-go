package arrayandhashing

func isValidSudoku(board [][]byte) bool {
	rowMap := make([]map[byte]bool, 9)
	initializeMaps(rowMap)
	columnMap := make([]map[byte]bool, 9)
	initializeMaps(columnMap)
	boardMap := make([][]map[byte]bool, 3)
	for rIndex := range boardMap {
		for i := 0; i < 3; i++ {
			boardMap[rIndex] = make([]map[byte]bool, 3)
		}
		initializeMaps(boardMap[rIndex])
	}
	for row, rowVal := range board {
		for col, colVal := range rowVal {
			if colVal == '.' {
				continue
			}
			rIndexForGrid, cIndexForGrid := row/3, col/3
			if rowMap[row][colVal] || columnMap[col][colVal] || boardMap[rIndexForGrid][cIndexForGrid][colVal] {
				return false
			}
			rowMap[row][colVal] = true
			columnMap[col][colVal] = true
			boardMap[rIndexForGrid][cIndexForGrid][colVal] = true
		}
	}
	return true
}

func initializeMaps(arr []map[byte]bool) {
	for index := range arr {
		arr[index] = make(map[byte]bool)
	}
}
