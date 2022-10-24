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

/**
I'm going to code out minimax here without using any help from any website and with the rough understanding of
Game theory. So we as the AI have to "maximize" our chances to win whereas the person on the other side wants to
"minimize" our chance to win. We have a selection of states to choose from, and we have to start at the first level
of states and go all the way down the tree of possibilities figuring out what is the end goal, then we go up the call
stack maximizing when we are on a level to choose from and minimizing when we are simulating the person, at the end each
of the states carry with themselves a score. We will choose the state with the highest score.
The way scoring works for a leaf node is as follows:
1. Leaf Node that wins = +1
2. Leaf Node that is a draw = 0
3. Leaf Node that is a lose = -1

we need a function that takes a board and gives us the possible moves we can make
for each of the candidate moves we create a branch where we simulate the opposite player
making a move or if that is the end state we return a score from the scoring criteria listed above.
When we start doing returns we will do a return which will also carry who the player who made the move is
if it is the opposition who made the last move we will maximize the score, if it was us who made the last move
and we do a return up the stack we minimize the score (act as a rational opposition). We do this all the way up
and use the last returned choice as our row and col to fill with the "O".
**/
func (t *TicTacToe) AiMove() (string, error) {
	// use minimax algorithm to make a move and return check_win

	return "O", nil
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
