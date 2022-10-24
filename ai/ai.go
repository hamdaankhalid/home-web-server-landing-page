package ai

import "github.com/hamdaankhalid/home-web-server-landing-page/utils"

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
type BestMove struct {
	Row, Col int
	score    int
}

// pseudo-const
var valueFuncMapper = map[string]int{"O": 1, "D": 0, "X": -1}

// AI / Player O is maximizer
func Minimax(board [][]string) BestMove {
	allPossibleMoves := getAllPossibleMoves(board)
	best_move := BestMove{-1, -1, -100}
	for _, move := range allPossibleMoves {
		board[move[0]][move[1]] = "O"
		score := minimaxHelper(board, false)
		board[move[0]][move[1]] = "-"

		if best_move.score < score {
			best_move = BestMove{move[0], move[1], score}
		}
	}
	return best_move
}

func minimaxHelper(board [][]string, isMaximizer bool) int {
	// base case is evaluated by our value func
	win := utils.CheckWinHelper(board)
	if win != "" {
		return valueFuncMapper[win]
	}

	var candidate string
	// since we were the maximizer the next player move we are simulating
	// will be the opposite
	if isMaximizer {
		candidate = "O"
	} else {
		candidate = "X"
	}
	possible_moves := getAllPossibleMoves(board)
	scores := make([]int, len(possible_moves))

	for i, move := range possible_moves {
		board[move[0]][move[1]] = candidate
		scores[i] = minimaxHelper(board, !isMaximizer)
		board[move[0]][move[1]] = "-"
	}

	const MaxUint = ^uint(0)
	if isMaximizer {
		max_e := -int(MaxUint >> 1)
		for _, e := range scores {
			if e > max_e {
				max_e = e
			}
		}
		return max_e
	} else {
		min_e := int(MaxUint >> 1)
		for _, e := range scores {
			if e < min_e {
				min_e = e
			}
		}
		return min_e
	}
}

func getAllPossibleMoves(board [][]string) [][]int {
	possibleMoves := [][]int{}
	for i, row := range board {
		for j, cell := range row {
			if cell == "-" {
				possibleMoves = append(possibleMoves, []int{i, j})
			}
		}
	}
	return possibleMoves
}
