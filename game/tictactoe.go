package game

import "errors"

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
		return "O"
	}
}

func (t *TicTacToe) Move(row int, col int) (string, error) {
	if row < 0 || row > 2 || col < 0 || col > 2 || t.board[row][col] != "-" {
		return "", errors.New("bad value")
	}

	if t.is_x_turn {
		t.board[row][col] = "X"
	} else {
		t.board[row][col] = "O"
	}

	allfull := true
	for _, row := range t.board {
		for _, cell := range row {
			if cell == "-" {
				allfull = false
			}
		}
	}

	win := t.check_win(row, col)
	if win && t.is_x_turn {
		return "X", nil
	} else if win && !t.is_x_turn {
		return "O", nil
	}

	if allfull {
		return "D", nil
	}

	t.is_x_turn = !t.is_x_turn
	return "", nil
}

func (t *TicTacToe) GetBoard() [][]string {
	return t.board
}

func (t *TicTacToe) check_win(row int, col int) bool {
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
