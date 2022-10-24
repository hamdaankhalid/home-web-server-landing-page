package utils

func CheckWinHelper(board [][]string) string {
	players := []string{"X", "O"}
	for _, player := range players {
		for i := 0; i < 3; i++ {
			if checkHorizontal(i, player, board) || checkVertical(i, player, board) {
				return player
			}
		}
		if checkTopLRDiag(player, board) {
			return player
		}
		if checkTopRLDiag(player, board) {
			return player
		}
	}

	for _, row := range board {
		for _, cell := range row {
			if cell == "-" {
				return ""
			}
		}
	}
	// if we didn't return from the block above we have a draw
	return "D"
}

func checkHorizontal(row int, candidate string, board [][]string) bool {
	for i := 0; i < 3; i++ {
		if board[row][i] != candidate {
			return false
		}
	}
	return true
}

func checkVertical(col int, candidate string, board [][]string) bool {
	for i := 0; i < 3; i++ {
		if board[i][col] != candidate {
			return false
		}
	}
	return true
}

func checkTopLRDiag(candidate string, board [][]string) bool {
	return board[0][0] == candidate && board[1][1] == candidate && board[2][2] == candidate
}

func checkTopRLDiag(candidate string, board [][]string) bool {
	return board[0][2] == candidate && board[1][1] == candidate && board[2][0] == candidate
}
