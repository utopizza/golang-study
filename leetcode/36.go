package leetcode

func isValidSudoku(board [][]byte) bool {
	if len(board) != 9 || len(board[0]) != 9 {
		return false
	}
	return validRows(board) && validColumns(board) && validSubCubes(board)
}

func validRows(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		set := make(map[byte]bool)
		for j := 0; j < 9; j++ {
			x := board[i][j]
			if x == '.' {
				continue
			}
			if set[x] {
				return false
			}
			set[x] = true
		}
	}
	return true
}

func validColumns(board [][]byte) bool {
	for j := 0; j < 9; j++ {
		set := make(map[byte]bool)
		for i := 0; i < 9; i++ {
			x := board[i][j]
			if x == '.' {
				continue
			}
			if set[x] {
				return false
			}
			set[x] = true
		}
	}
	return true
}

func validSubCubes(board [][]byte) bool {
	var sets [3][3]map[byte]bool
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sets[i][j] = make(map[byte]bool)
		}
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			x := board[i][j]
			if x == '.' {
				continue
			}
			if sets[i/3][j/3][x] {
				return false
			}
			sets[i/3][j/3][x] = true
		}
	}
	return true
}
