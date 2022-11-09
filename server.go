package main

import (
	"fmt"
	"github.com/hamdaankhalid/home-web-server-landing-page/dbclient"
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
	dbClient   *dbclient.Db
}

func initServer() *serverState {
	dbClient := dbclient.Db{DbHost: "localhost", Port: "8000"}

	xScore, err := dbClient.GetInt("X")
	if err != nil {
		xScore = 0
	}
	oScore, err := dbClient.GetInt("O")
	if err != nil {
		oScore = 0
	}
	dScore, err := dbClient.GetInt("D")
	if err != nil {
		dScore = 0
	}
	initMap := map[string]int{"X": xScore, "O": oScore, "D": dScore}

	return &serverState{game: game.InitTicTacToe(), scoreTable: initMap, dbClient: &dbClient}
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
		s.dbClient.SetInt(win, s.scoreTable[win]+1)
		s.scoreTable[win] += 1
		return win
	}
	return ""
}

// LandingPageData all members are public since this is used in a different package
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
			currBoard := state.game.GetBoard()
			currPlayer := state.game.GetPlayerTurn()
			tmplArgs := LandingPageData{TimeNow: time.Now().String(), Board: currBoard, Xscore: state.scoreTable["X"], Oscore: state.scoreTable["O"], Drawscore: state.scoreTable["D"], PlayerTurn: currPlayer, History: state.game.History}
			tmpl.Execute(w, tmplArgs)
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
