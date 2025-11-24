package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Game struct {
	Grid     [6][7]string
	Turn     string
	Winner   string
	IsDraw   bool
	Message  string
	Player1  string
	Player2  string
	Color1   string
	Color2   string
	Turns    int
}

type GameHistory struct {
	Player1 string
	Player2 string
	Winner  string
	Date    string
	Turns   int
}

var game Game
var scoreboard []GameHistory

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/game/init", initHandler)
	http.HandleFunc("/game/init/traitement", initTraitementHandler)
	http.HandleFunc("/game/play", playHandler)
	http.HandleFunc("/game/play/traitement", playTraitementHandler)
	http.HandleFunc("/game/end", endHandler)
	http.HandleFunc("/game/scoreboard", scoreboardHandler)
	http.HandleFunc("/reset", resetHandler)

	http.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir("statics"))))

	log.Println("Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func resetGame() {
	for i := range game.Grid {
		for j := range game.Grid[i] {
			game.Grid[i][j] = "empty"
		}
	}
	game.Winner = ""
	game.IsDraw = false
	game.Message = ""
	game.Turns = 0
	game.Turn = game.Color1
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	_ = tmpl.Execute(w, nil)
}

func initHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/init.html"))
	_ = tmpl.Execute(w, nil)
}

func initTraitementHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	player1 := r.FormValue("player1")
	player2 := r.FormValue("player2")
	color1 := r.FormValue("color1")
	color2 := r.FormValue("color2")

	if color1 == color2 {
		http.Error(w, "Les deux joueurs ne peuvent pas avoir la même couleur", http.StatusBadRequest)
		return
	}

	game.Player1 = player1
	game.Player2 = player2
	game.Color1 = color1
	game.Color2 = color2
	game.Turn = color1
	resetGame()

	http.Redirect(w, r, "/game/play", http.StatusSeeOther)
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/play.html"))
	err := tmpl.Execute(w, game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func playTraitementHandler(w http.ResponseWriter, r *http.Request) {
	if game.Winner != "" || game.IsDraw {
		http.Redirect(w, r, "/game/end", http.StatusSeeOther)
		return
	}

	colStr := r.URL.Query().Get("col")
	col, err := strconv.Atoi(colStr)
	if err != nil || col < 0 || col > 6 {
		http.Redirect(w, r, "/game/play", http.StatusSeeOther)
		return
	}

	for row := 5; row >= 0; row-- {
		if game.Grid[row][col] == "empty" {
			game.Grid[row][col] = game.Turn
			game.Turns++

			if checkWin(row, col) {
				if game.Turn == game.Color1 {
					game.Message = game.Player1 + " a gagné 🎉"
				} else {
					game.Message = game.Player2 + " a gagné 🎉"
				}
				game.Winner = game.Turn
				scoreboard = append(scoreboard, GameHistory{
					Player1: game.Player1,
					Player2: game.Player2,
					Winner:  game.Message,
					Date:    time.Now().Format("02/01/2006 15:04"),
					Turns:   game.Turns,
				})
				http.Redirect(w, r, "/game/end", http.StatusSeeOther)
				return
			} else if checkDraw() {
				game.Message = "Match nul 🤝"
				game.IsDraw = true
				scoreboard = append(scoreboard, GameHistory{
					Player1: game.Player1,
					Player2: game.Player2,
					Winner:  "Égalité",
					Date:    time.Now().Format("02/01/2006 15:04"),
					Turns:   game.Turns,
				})
				http.Redirect(w, r, "/game/end", http.StatusSeeOther)
				return
			} else {
				if game.Turn == game.Color1 {
					game.Turn = game.Color2
				} else {
					game.Turn = game.Color1
				}
			}
			break
		}
	}
	http.Redirect(w, r, "/game/play", http.StatusSeeOther)
}

func checkWin(row, col int) bool {
	color := game.Grid[row][col]
	directions := [][2]int{{0, 1}, {1, 0}, {1, 1}, {1, -1}}

	for _, d := range directions {
		count := 1
		for i := 1; i < 4; i++ {
			r, c := row+i*d[0], col+i*d[1]
			if r < 0 || r > 5 || c < 0 || c > 6 || game.Grid[r][c] != color {
				break
			}
			count++
		}
		for i := 1; i < 4; i++ {
			r, c := row-i*d[0], col-i*d[1]
			if r < 0 || r > 5 || c < 0 || c > 6 || game.Grid[r][c] != color {
				break
			}
			count++
		}
		if count >= 4 {
			return true
		}
	}
	return false
}

func checkDraw() bool {
	for _, row := range game.Grid {
		for _, cell := range row {
			if cell == "empty" {
				return false
			}
		}
	}
	return true
}

func endHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/end.html"))
	_ = tmpl.Execute(w, game)
}

func scoreboardHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/scoreboard.html"))
	_ = tmpl.Execute(w, scoreboard)
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	resetGame()
	http.Redirect(w, r, "/game/play", http.StatusSeeOther)
}