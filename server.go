package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/hamdaankhalid/home-web-server-landing-page/game"
	"github.com/hamdaankhalid/home-web-server-landing-page/static"
)

// no use of this struct outside this package
type serverState struct {
	game       *game.TicTacToe
	scoreTable map[string]int
	mu         sync.Mutex
}

func initServer() *serverState {
	init_map := map[string]int{"X": 0, "O": 0, "D": 0}
	return &serverState{game: game.InitTicTacToe(), scoreTable: init_map}
}

func (s *serverState) move(row int, col int) string {
	win, err := s.game.Move(row, col)
	if err != nil {
		return ""
	}

	if win == "" && s.game.GetPlayerTurn() == "O" {
		win = s.game.AiMove()
	}

	if win != "" {
		s.game = game.InitTicTacToe()
		s.scoreTable[win] += 1
		return win
	}
	return ""
}

// all members are public since this is used in a different package
type LandingPageData struct {
	TimeNow                   string
	Board                     [][]string
	Xscore, Oscore, Drawscore int
	PlayerTurn                string
	History                   []string
}

func actions(route string, state *serverState) func(w http.ResponseWriter, r *http.Request) {
	switch route {
	case "move":
		return func(w http.ResponseWriter, r *http.Request) {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			rval, err := strconv.Atoi(r.Form.Get("row"))
			if err != nil {
				http.Redirect(w, r, "/", http.StatusBadRequest)
				return
			}
			cval, err := strconv.Atoi(r.Form.Get("col"))
			if err != nil {
				http.Redirect(w, r, "/", http.StatusBadRequest)
				return
			}
			state.mu.Lock()
			defer state.mu.Unlock()
			state.move(rval, cval)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	default:
		return func(w http.ResponseWriter, r *http.Request) {
			state.mu.Lock()
			defer state.mu.Unlock()
			tmpl := template.Must(template.ParseFS(static.Assets, "landing-page.tmpl"))
			curr_board := state.game.GetBoard()
			curr_player := state.game.GetPlayerTurn()
			tmpl_args := LandingPageData{TimeNow: time.Now().String(), Board: curr_board, Xscore: state.scoreTable["X"], Oscore: state.scoreTable["O"], Drawscore: state.scoreTable["D"], PlayerTurn: curr_player, History: state.game.History}
			tmpl.Execute(w, tmpl_args)
		}
	}
}

func main() {
	state := initServer()

	http.HandleFunc("/tic-tac-toe-move", actions("move", state))
	http.HandleFunc("/", actions("", state))

	fmt.Printf("Started server at port 3001\n")
	if err := http.ListenAndServe(":3001", nil); err != nil {
		log.Fatal(err)
	}
}
