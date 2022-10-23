package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/hamdaankhalid/home-web-server-landing-page/game"
	"github.com/hamdaankhalid/home-web-server-landing-page/static"
)

type ServerState struct {
	Game       *game.TicTacToe
	ScoreTable map[string]int
}

func InitServer() *ServerState {
	init_map := map[string]int{"X": 0, "O": 0, "D": 0}
	return &ServerState{game.InitTicTacToe(), init_map}
}

func (s *ServerState) move(row int, col int) string {
	win, err := s.Game.Move(row, col)
	if err != nil {
		return ""
	}

	if win != "" {
		s.Game = game.InitTicTacToe()
		s.ScoreTable[win] += 1
		return win
	}
	return ""
}

type LandingPageData struct {
	TimeNow                   string
	Board                     [][]string
	Xscore, Oscore, Drawscore int
	PlayerTurn                string
}

func actions(route string, state *ServerState) func(w http.ResponseWriter, r *http.Request) {
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

			state.move(rval, cval)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	default:
		return func(w http.ResponseWriter, r *http.Request) {

			tmpl := template.Must(template.ParseFS(static.Assets, "landing-page.tmpl"))
			curr_board := state.Game.GetBoard()
			curr_player := state.Game.GetPlayerTurn()
			tmpl_args := LandingPageData{TimeNow: time.Now().String(), Board: curr_board, Xscore: state.ScoreTable["X"], Oscore: state.ScoreTable["O"], Drawscore: state.ScoreTable["D"], PlayerTurn: curr_player}
			tmpl.Execute(w, tmpl_args)
		}
	}
}

func main() {
	state := InitServer()

	http.HandleFunc("/tic-tac-toe-move", actions("move", state))
	http.HandleFunc("/", actions("", state))

	fmt.Printf("Started server at port 3001\n")
	if err := http.ListenAndServe(":3001", nil); err != nil {
		log.Fatal(err)
	}
}
