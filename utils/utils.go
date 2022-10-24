package utils

func CheckWinHelper(board [][]string) string {
	players := []string{"X", "O"}
	for _, player := range players {
		for i := 0; i < len(board); i++ {
			if check_horizontal_row(i, player, board) {
				return player
			}
		}

		for i := 0; i < len(board[0]); i++ {
			if check_vertical_col(i, player, board) {
				return player
			}
		}

		if check_top_l_r_diag(player, board) {
			return player
		}

		if check_top_r_l_diag(player, board) {
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

func check_horizontal_row(row int, candidate string, board [][]string) bool {
	for i := 0; i < len(board[0]); i++ {
		if board[row][i] != candidate {
			return false
		}
	}
	return true
}

func check_vertical_col(col int, candidate string, board [][]string) bool {
	for i := 0; i < len(board); i++ {
		if board[i][col] != candidate {
			return false
		}
	}
	return true
}

func check_top_l_r_diag(candidate string, board [][]string) bool {
	return board[0][0] == candidate && board[1][1] == candidate && board[2][2] == candidate
}

func check_top_r_l_diag(candidate string, board [][]string) bool {
	return board[0][2] == candidate && board[1][1] == candidate && board[2][0] == candidate
}
