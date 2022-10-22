package game

type TicTacToe struct {
	board     [][]string
	is_x_turn bool
}

func InitTicTacToe() *TicTacToe {
	init_board := [][]string{{"-", "-", "-"}, {"-", "-", "-"}, {"-", "-", "-"}}
	t := TicTacToe{init_board, true}
	return &t
}

func (t *TicTacToe) GetPlayerTurn() string {
	if t.is_x_turn {
		return "X"
	} else {
		return "Y"
	}
}

func (t *TicTacToe) Move(row int, col int) string {
	if t.is_x_turn {
		t.board[row][col] = "X"
	} else {
		t.board[row][col] = "O"
	}

	win := t.check_win(row, col)
	if win && t.is_x_turn {
		return "X"
	} else if win && !t.is_x_turn {
		return "O"
	}

	t.is_x_turn = !t.is_x_turn
	return ""
}

func (t *TicTacToe) GetBoard() [][]string {
	return t.board
}

func (t *TicTacToe) check_win(row int, col int) bool {
	// make a move and check if won
	var candidate string
	if t.is_x_turn {
		candidate = "X"
	} else {
		candidate = "O"
	}

	horizontal_win := t.board[row][0] == candidate && t.board[row][1] == candidate && t.board[row][2] == candidate
	vertical_win := t.board[0][col] == candidate && t.board[1][col] == candidate && t.board[2][col] == candidate
	diagonal_win_1 := t.board[0][0] == candidate && t.board[1][1] == candidate && t.board[2][2] == candidate
	diagonal_win_2 := t.board[2][0] == candidate && t.board[1][1] == candidate && t.board[0][2] == candidate
	return horizontal_win || vertical_win || diagonal_win_1 || diagonal_win_2
}
