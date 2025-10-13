package power4

import (
	"log"
	"os"
	"sync"
)

type Game struct {
	Rows, Cols    int
	Grid          [][]int // 0 vide, 1 P1, 2 P2
	Player1       string
	Player2       string
	Current       int
	Moves         int
	Winner        int // 0=none, 1|2=gagnant, 3=nul
	Difficulty    string
	Started       bool
	ColumnIndexes []int
}

var (
	game Game
	mu   sync.Mutex
)

func seq(n int) []int {
	out := make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = i
	}
	return out
}

func newGame(diff, p1, p2 string) Game {
	var rows, cols int
	switch diff {
	case "easy":
		rows, cols = 6, 7
	case "normal":
		rows, cols = 6, 9
	case "hard":
		rows, cols = 7, 8
	default:
		rows, cols = 6, 7
	}
	grid := make([][]int, rows)
	for r := 0; r < rows; r++ {
		grid[r] = make([]int, cols)
	}
	return Game{
		Rows: rows, Cols: cols, Grid: grid,
		Player1: p1, Player2: p2,
		Current:       1,
		Difficulty:    diff,
		Started:       true,
		ColumnIndexes: seq(cols),
	}
}

func dropToken(g *Game, col int, player int) bool {
	for r := g.Rows - 1; r >= 0; r-- {
		if g.Grid[r][col] == 0 {
			g.Grid[r][col] = player
			return true
		}
	}
	return false
}

func checkWin(grid [][]int, rows, cols, p int) bool {
	// horizontal
	for r := 0; r < rows; r++ {
		for c := 0; c <= cols-4; c++ {
			if grid[r][c] == p && grid[r][c+1] == p && grid[r][c+2] == p && grid[r][c+3] == p {
				return true
			}
		}
	}
	// vertical
	for c := 0; c < cols; c++ {
		for r := 0; r <= rows-4; r++ {
			if grid[r][c] == p && grid[r+1][c] == p && grid[r+2][c] == p && grid[r+3][c] == p {
				return true
			}
		}
	}
	// diag ↘
	for r := 0; r <= rows-4; r++ {
		for c := 0; c <= cols-4; c++ {
			if grid[r][c] == p && grid[r+1][c+1] == p && grid[r+2][c+2] == p && grid[r+3][c+3] == p {
				return true
			}
		}
	}
	// diag ↗
	for r := 3; r < rows; r++ {
		for c := 0; c <= cols-4; c++ {
			if grid[r][c] == p && grid[r-1][c+1] == p && grid[r-2][c+2] == p && grid[r-3][c+3] == p {
				return true
			}
		}
	}
	return false
}

func mustLogPaths() {
	paths := []string{
		"templates/home.html",
		"templates/play.html",
		"templates/result.html",
		"static/style.css",
	}
	for _, p := range paths {
		if _, err := os.Stat(p); err != nil {
			log.Printf("ABSENT: %s -> %v\n", p, err)
		} else {
			log.Printf("OK     : %s\n", p)
		}
	}
}