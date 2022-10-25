package game

import (
	"errors"
	"strconv"

	"github.com/hamdaankhalid/home-web-server-landing-page/ai"
	"github.com/hamdaankhalid/home-web-server-landing-page/utils"
)

type TicTacToe struct {
	board   [][]string
	isXTurn bool
	History []string
}

func InitTicTacToe() *TicTacToe {
	init_board := [][]string{{"-", "-", "-"}, {"-", "-", "-"}, {"-", "-", "-"}}
	history := []string{}
	t := TicTacToe{init_board, true, history}
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

	t.board[row][col] = "X"

	t.History = append(t.History, "Human placed X at "+strconv.Itoa(row)+", "+strconv.Itoa(col))

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

	t.History = append(t.History, "AI placed O at "+strconv.Itoa(result.Row)+", "+strconv.Itoa(result.Col))

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
