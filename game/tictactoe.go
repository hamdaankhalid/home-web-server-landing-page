package game

import (
	"errors"

	"github.com/hamdaankhalid/home-web-server-landing-page/ai"
	"github.com/hamdaankhalid/home-web-server-landing-page/utils"
)

type TicTacToe struct {
	board   [][]string
	isXTurn bool
}

func InitTicTacToe() *TicTacToe {
	init_board := [][]string{{"-", "-", "-"}, {"-", "-", "-"}, {"-", "-", "-"}}
	t := TicTacToe{init_board, true}
	return &t
}

func (t *TicTacToe) GetPlayerTurn() string {
	if t.isXTurn {
		return "X"
	} else {
		return "O"
	}
}

func (t *TicTacToe) Move(row int, col int) (string, error) {
	if row < 0 || row > 2 || col < 0 || col > 2 || t.board[row][col] != "-" {
		return "", errors.New("bad value")
	}

	if t.isXTurn {
		t.board[row][col] = "X"
	} else {
		t.board[row][col] = "O"
	}

	win := utils.CheckWinHelper(t.board)
	if win != "" {
		return win, nil
	}

	t.isXTurn = !t.isXTurn
	return "", nil
}

func (t *TicTacToe) AiMove() string {
	// use minimax algorithm to make a move and return check_win
	result := ai.Minimax(t.board)
	t.board[result.Row][result.Col] = "O"

	win := utils.CheckWinHelper(t.board)
	if win != "" {
		return win
	}

	t.isXTurn = !t.isXTurn
	return ""
}

func (t *TicTacToe) GetBoard() [][]string {
	return t.board
}
